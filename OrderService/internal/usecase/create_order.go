package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
)

func (o *OrderService) CreateOrder(ctx context.Context, newOrder order.Order) (int64, string, error) {

	marketsID, err := o.GetEnableMarkets(ctx)
	if err != nil {
		return 0, "", err
	}

	//In-Memory реализация
	orderID, err := o.AddOrderID(newOrder, marketsID)
	if err != nil {
		return 0, "", err
	}
	status, err := o.AddOrderStorage(newOrder, orderID)
	if err != nil {
		return 0, "", err
	}

	return int64(orderID), status, nil
}
