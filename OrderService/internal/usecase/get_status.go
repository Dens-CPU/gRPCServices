package usecase

import "Academy/gRPCServices/OrderService/internal/domain/order"

func (o *OrderService) GetStatus(key order.Key) (string, error) {
	status, err := o.GetOrderState(key)
	if err != nil {
		return "", err
	}
	return status, nil
}
