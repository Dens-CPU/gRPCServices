package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
)

type Storage interface {
	AddOrderStorage(context.Context, order.Order, []int64) (string, string, error) //Добавление нового заказа в хранилище
	GetOrderState(context.Context, order.Key) (string, error)                      //Получение статуса заказа
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
