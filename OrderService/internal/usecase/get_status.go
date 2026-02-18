package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
)

func (o *OrderService) GetStatus(ctx context.Context, key order.Key) (string, error) {
	status, err := o.GetOrderState(ctx, key)
	if err != nil {
		return "", err
	}
	return status, nil
}
