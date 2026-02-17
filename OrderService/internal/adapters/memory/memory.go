package memory

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"sync"
)

// In-memory хранилище
type Storage struct {
	date     map[order.Key]*order.Order
	ordersID map[int]struct{}
	mu       sync.Mutex
}

// Конструктор для In-memory
func NewStorage() *Storage {
	return &Storage{date: make(map[order.Key]*order.Order), ordersID: make(map[int]struct{})}
}
