package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
)

type Storage interface {
	AddOrderStorage(order.Order, int) (string, error) //Добавление нового заказа в хранилище
	GetOrderState(order.Key) (string, error)          //Получение статуса заказа
	AddOrderID(order.Order, []int64) (int, error)     //Создание ID заказа
}

type MarketsService interface {
	GetEnableMarkets(context.Context) ([]int64, error) //Получение списка доступных рынков
}

type OrderService struct {
	Storage
	MarketsService
}

func NewOrderServ(in_memory Storage, markets_service MarketsService) *OrderService {
	return &OrderService{in_memory, markets_service}
}
