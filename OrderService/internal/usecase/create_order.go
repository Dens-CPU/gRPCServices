package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"

	"go.uber.org/zap"
)

func (o *OrderService) CreateOrder(ctx context.Context, newOrder order.Order) (string, string, error) {

	tracer := o.trace.Tracer("OrderService")

	//Получения списка доступных рынков
	ctx, span := tracer.Start(ctx, "Get enable markets")
	defer span.End()
	marketsID, err := o.GetEnableMarkets(ctx)
	if err != nil {
		o.logger.Error("ошибка получения доступныхх рынков:",
			zap.Error(err),
		)
		return "", "", err
	}
	o.logger.Info("Получен список доступных рынков",
		zap.String("get enable markets span ID:", span.SpanContext().TraceID().String()),
	)

	//Создание нового заказа
	ctx, span = tracer.Start(ctx, "AddOrder")
	defer span.End()
	orderID, status, err := o.AddOrderStorage(ctx, newOrder, marketsID)
	if err != nil {
		return "", "", err
	}
	o.logger.Info("Создан заказ:",
		zap.String("OrderID:", orderID),
		zap.String("add order span ID:", span.SpanContext().TraceID().String()),
	)

	//Выполнение заказа
	stateCh := o.ControlOrder(newOrder.Order_type, newOrder.User_id, orderID)

	go o.Notify.AddNewState(newOrder.User_id, orderID, stateCh)

	return orderID, status, nil
}
