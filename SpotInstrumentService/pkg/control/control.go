// В данном файле реализована проверка достпности сайтов
package control

import (
	"Academy/gRPCServices/SpotInstrumentService/pkg/memory"
	"context"
	"math/rand"
	"sync"

	"fmt"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Управленеи рынками
func AccessControl(markets []string, storage map[string]*memory.Market) string {
	var mu sync.Mutex
	//Заполнение storage
	for i, m := range markets {
		mu.Lock()
		storage[m] = &memory.Market{ID: int64(i), Name: m, Enable: true, Delete_at: nil}
		mu.Unlock()
	}

	//Создание контекста для времени управления состояниями рынков
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Время жизни цикла истекло")
			return "Управленение работой ранков завершено"
		default:
			d := rand.Intn(3) //Случайная блокировка или удаление рынка
			switch d {

			case 0: //Блокировка доступа случайного маркета

				n := rand.Intn(len(markets))

				mu.Lock()
				key := markets[n]
				if storage[key].Enable != false { //Проверка доступа к рынку
					storage[key].Enable = false
					mu.Unlock()
					break
				}
				mu.Unlock()

			case 1: //Удаление случайного маркета с рынка
				n := rand.Intn(len(markets))

				mu.Lock()
				key := markets[n]
				if storage[key].Enable != false { //Проверка доступа к рынку
					storage[key].Enable = false
					delete_at := time.Now().Local()
					storage[key].Delete_at = &delete_at
					mu.Unlock()
					break
				}
				mu.Unlock()

			case 2: //Востановление доступа к маркету на рынке

				for { //Поиск маркета, который недоступен
					n := rand.Intn(len(markets))

					mu.Lock()
					key := markets[n]
					if storage[key].Enable == false {
						storage[key].Enable = true
						storage[key].Delete_at = nil
						mu.Unlock()
						break
					}
					mu.Unlock()
				}

			}
		}
		time.Sleep(10 * time.Second)
	}
}
