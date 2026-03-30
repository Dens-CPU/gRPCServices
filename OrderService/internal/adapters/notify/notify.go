package notify

import (
	"sync"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
)

type StatusStorage struct {
	Status map[orderdomain.Key]string
	Subs   map[orderdomain.Key][]chan string
	mu     sync.Mutex
}

func NewStatStorage() *StatusStorage {
	return &StatusStorage{Status: make(map[orderdomain.Key]string), Subs: make(map[orderdomain.Key][]chan string)}
}

func (s *StatusStorage) GetNumbersSubsChan(key orderdomain.Key) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	l := len(s.Subs[key])
	return l
}
