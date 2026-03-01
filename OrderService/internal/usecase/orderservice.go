package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
)

type Storage interface {
	AddOrderStorage(context.Context, order.Order, []int64) (string, string, error) //Добавление нового заказа в хранилище
	GetOrderState(context.Context, order.Key) (string, error)                      //Получение статуса заказа
	ControlOrder(string, int64, string) chan string
}

type MarketsService interface {
	GetEnableMarkets(context.Context) ([]int64, error) //Получение списка доступных рынков
}

type Notify interface {
	AddNewState(int64, string, chan string)
	GetStatus(order.Key) string
	AddNewSub(order.Key) chan string
	GetNumbersSubsChan(order.Key) int
	UpdateStatusSubs(order.Key)
}

type OrderService struct {
	Storage
	MarketsService
	Notify
}

func NewOrderServ(in_memory Storage, markets_service MarketsService, notify Notify) *OrderService {
	return &OrderService{in_memory, markets_service, notify}
}
