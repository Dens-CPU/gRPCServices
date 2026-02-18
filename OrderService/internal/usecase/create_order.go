package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
)

func (o *OrderService) CreateOrder(ctx context.Context, newOrder order.Order) (string, string, error) {

	marketsID, err := o.GetEnableMarkets(ctx)
	if err != nil {
		return "", "", err
	}

	orderID, status, err := o.AddOrderStorage(ctx, newOrder, marketsID)
	if err != nil {
		return "", "", err
	}

	return orderID, status, nil
}
