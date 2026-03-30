package notify

import orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"

// Получение актуального статуса заказа
func (s *StatusStorage) GetStatus(key orderdomain.Key) string {
	s.mu.Lock()
	status := s.Status[key]
	s.mu.Unlock()
	if status == "" {
		return "created"
	}
	return status
}
