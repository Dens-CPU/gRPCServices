package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"

	"go.uber.org/zap"
)

func (o *OrderService) GetStatus(ctx context.Context, key order.Key) (string, error) {
	tracer := o.trace.Tracer("OrderService")

	ctx, span := tracer.Start(ctx, "Get status")
	defer span.End()
	status, err := o.GetOrderState(ctx, key)
	if err != nil {
		o.logger.Error("ошибка получения статуса заказа:",
			zap.Error(err),
		)
		return "", err
	}
	return status, nil
}
