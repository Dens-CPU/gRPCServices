package notify

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"fmt"
	"time"
)

func (s *StatusStorage) UpdateStatusSubs(key order.Key) {
	go func() {
		var laststatus string
		for {
			status := s.GetStatus(key)
			if laststatus != status {
				// Рассылаем всем подписчикам
				for _, ch := range s.Subs[key] {
					select {
					case ch <- status:
					default:
					}
				}
			} else {
				for _, ch := range s.Subs[key] {
					close(ch)
					fmt.Println("Канал закрыт")
				}
				return
			}
			laststatus = status
			time.Sleep(3 * time.Second)
		}
	}()
}
