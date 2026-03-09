package notify

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
)

// Добавление нового статуса заказа
func (s *StatusStorage) AddNewState(user_id int64, orderId string, statCh chan string) {
	key := order.Key{User_id: user_id, Order_id: orderId}
	for state := range statCh {
		s.mu.Lock()
		s.Status[key] = state

		s.mu.Unlock()
	}
}
