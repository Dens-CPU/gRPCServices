package memory

import (
	ordererrors "Academy/gRPCServices/OrderService/internal/domain/error"
	"Academy/gRPCServices/OrderService/internal/domain/order"
)

func (s *Storage) GetOrderState(key order.Key) (string, error) {
	if _, exist := s.date[key]; !exist {
		return "", ordererrors.Not_exist_order
	}
	status := s.date[key].Status
	return status, nil
}
