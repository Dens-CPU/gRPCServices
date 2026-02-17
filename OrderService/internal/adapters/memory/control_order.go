package memory

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"strings"
	"sync"
	"time"
)

// Контроль выполнения заказов
func (s *Storage) ControlStat(key order.Key) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		orderType := strings.ToLower(s.date[key].Order_type)
		switch orderType {

		case "normal":
			time.Sleep(5 * time.Second)
			s.mu.Lock()
			s.date[key].Status = "in progress"
			s.mu.Unlock()
			time.Sleep(5 * time.Second)
			s.mu.Lock()
			s.date[key].Status = "complpeted"
			s.mu.Unlock()

		case "express":
			time.Sleep(2 * time.Second)
			s.mu.Lock()
			s.date[key].Status = "in progress"
			s.mu.Unlock()
			time.Sleep(2 * time.Second)
			s.mu.Lock()
			s.date[key].Status = "complpeted"
			s.mu.Unlock()
		}
	}()
	go func() {
		wg.Wait()
	}()
}
