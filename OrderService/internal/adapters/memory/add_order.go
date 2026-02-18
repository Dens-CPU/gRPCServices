package memory

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
)

// Добавленеи заказа
func (s *Storage) AddOrderStorage(newOrder order.Order, orderID string, marketID []int64) (string, error) {

	orderID, err := s.AddOrderID(newOrder, marketID)
	if err != nil {
		return "", err
	}

	key := order.Key{User_id: newOrder.User_id, Order_id: orderID} //Создание ключа для in-memory хранилища
	s.date[key] = &newOrder
	s.date[key].Status = "created" //Сохранение заказа в памяти

	s.ControlStat(key)
	return "created", nil
}
