package ordermemory

import (
	"errors"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Структура закакза
type Order struct {
	User_id    int64
	Market_id  int64
	Order_type string
	Price      float64
	Quantity   int64
	Status     string
}

type Key struct {
	User_id  int64
	Order_id int
}

// In-memory хранилище
type Storage struct {
	date     map[Key]*Order
	ordersID map[int]struct{}
	mu       sync.Mutex
}

// Хранилище индетификаторов заказов
type OrdersID struct {
}

// Конструктор для In-memory
func NewStorage() *Storage {
	return &Storage{date: make(map[Key]*Order), ordersID: make(map[int]struct{})}
}

// Добавленеи заказа
func (s *Storage) AddOrder(marketsID []int64, newOrder Order) (int64, string, error) {

	var foundMarket bool //Флаг, показывающий найден нужный рынок или нет

	for _, mId := range marketsID { //Проверка наличия нужного рынка
		if mId == newOrder.Market_id {
			foundMarket = true
		}
	}
	if foundMarket != true {
		return 0, "", errors.New("Рынок недоступен")
	}

	var orderId int //ID нового заказа
	for {
		id := rand.Intn(math.MaxInt64) //Создание ID заказа
		if _, exist := s.ordersID[id]; !exist {
			s.ordersID[id] = struct{}{}
			orderId = id
			break
		}
	}

	key := Key{User_id: newOrder.User_id, Order_id: orderId} //Создание ключа для in-memory хранилища
	s.date[key] = &newOrder
	s.date[key].Status = "created" //Сохранение заказа в памяти

	s.ControlStat(key)
	return int64(orderId), "create", nil
}

func (s *Storage) GetOrderState(key Key) (string, error) {
	if _, exist := s.date[key]; !exist {
		return "", errors.New("Заказа не существует")
	}
	status := s.date[key].Status
	return status, nil
}

// Контроль выполнения заказов
func (s *Storage) ControlStat(key Key) {
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
