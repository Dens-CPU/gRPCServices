package controlorderstat

import (
	"Academy/gRPCServices/OrderService/pkg/memory"
	"strings"
	"sync"
	"time"
)

func ControlStat(key memory.Key, storage map[memory.Key]*memory.Order) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		orderType := strings.ToLower(storage[key].Order_type)
		switch orderType {

		case "normal":
			time.Sleep(5 * time.Second)
			mu.Lock()
			storage[key].Status = "in progress"
			mu.Unlock()
			time.Sleep(5 * time.Second)
			mu.Lock()
			storage[key].Status = "complpeted"
			mu.Unlock()

		case "express":
			time.Sleep(2 * time.Second)
			mu.Lock()
			storage[key].Status = "in progress"
			mu.Unlock()
			time.Sleep(2 * time.Second)
			mu.Lock()
			storage[key].Status = "complpeted"
			mu.Unlock()
		}
	}()
	go func() {
		wg.Wait()
	}()
}
