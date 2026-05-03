package usecase

import (
	"context"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	"go.uber.org/zap"
)

func (o *OrderService) CreateOrder(ctx context.Context, newOrder orderdomain.Order) (string, string, error) {
	//Get marketss list
	ctx, span := o.tracer.Start(ctx, "Enable markets")
	defer span.End()

	markets, err := o.GetEnableMarkets(ctx, newOrder.UserId, newOrder.UserRole)
	if err != nil {
		o.logger.Error("Error getting available markets:",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return "", "", err
	}
	o.logger.Info("Received a list of available markets",
		zap.String("spanID:", span.SpanContext().SpanID().String()),
	)

	//Create a new order
	ctx, span = o.tracer.Start(ctx, "Add order")
	defer span.End()

	orderID, status, err := o.AddOrderStorage(ctx, newOrder, markets)
	if err != nil {
		o.logger.Error("Error add new order to storage:",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return "", "", err
	}
	o.logger.Info("Order created:",
		zap.String("OrderID:", orderID),
		zap.String("spanID:", span.SpanContext().SpanID().String()),
	)

	//Order fulfillment
	o.logger.Info("The order has been processed")

	return orderID, status, nil
}
