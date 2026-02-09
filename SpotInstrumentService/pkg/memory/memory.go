// В данном файле прописаны используемые структуры для in-memory хралилища
package spotmemory

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

// Управление рынками
func init() {
	rand.Seed(time.Now().Unix())
}

// Структура Market
type Market struct {
	ID        int64
	Name      string
	Enable    bool
	Delete_at *time.Time
}

// Хранилище данных о рынках
type Storage struct {
	date map[string]*Market
	mu   sync.RWMutex
}

func NewStorage(size int) *Storage {
	return &Storage{date: make(map[string]*Market, size)}
}

// Получение доступных рынков
func (s *Storage) GetEnableMarkets() []*Market {
	var enableMarkets []*Market

	for _, value := range s.date {
		s.mu.RLock()
		if value.Delete_at == nil || value.Enable == true {
			enableMarkets = append(enableMarkets, value)
		}
		s.mu.RUnlock()
	}
	return enableMarkets
}

// Управленеи рынками
func (s *Storage) AccessControl(markets []string) string {
	//Заполнение storage
	for i, m := range markets {
		s.mu.Lock()
		s.date[m] = &Market{ID: int64(i), Name: m, Enable: true, Delete_at: nil}
		s.mu.Unlock()
	}

	//Создание контекста для времени управления состояниями рынков
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			// fmt.Println("Время жизни цикла истекло")
			return "Управленение работой рынков завершено"
		default:
			d := rand.Intn(3) //Случайная блокировка или удаление рынка
			switch d {

			case 0: //Блокировка доступа случайного маркета

				n := rand.Intn(len(markets))

				s.mu.Lock()
				key := markets[n]
				if s.date[key].Enable != false { //Проверка доступа к рынку
					s.date[key].Enable = false
					s.mu.Unlock()
					break
				}
				s.mu.Unlock()

			case 1: //Удаление случайного маркета с рынка
				n := rand.Intn(len(markets))

				s.mu.Lock()
				key := markets[n]
				if s.date[key].Enable != false { //Проверка доступа к рынку
					s.date[key].Enable = false
					delete_at := time.Now().Local()
					s.date[key].Delete_at = &delete_at
					s.mu.Unlock()
					break
				}
				s.mu.Unlock()

			case 2: //Востановление доступа к маркету на рынке

				for { //Поиск маркета, который недоступен
					n := rand.Intn(len(markets))

					s.mu.Lock()
					key := markets[n]
					if s.date[key].Enable == false {
						s.date[key].Enable = true
						s.date[key].Delete_at = nil
						s.mu.Unlock()
						break
					}
					s.mu.Unlock()
				}

			}
		}
		time.Sleep(1 * time.Second)
	}
}
