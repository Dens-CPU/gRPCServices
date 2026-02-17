package memory

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
)

// Добавленеи заказа
func (s *Storage) AddOrderStorage(newOrder order.Order, orderID int) (string, error) {

	key := order.Key{User_id: newOrder.User_id, Order_id: orderID} //Создание ключа для in-memory хранилища
	s.date[key] = &newOrder
	s.date[key].Status = "created" //Сохранение заказа в памяти

	s.ControlStat(key)
	return "created", nil
}
