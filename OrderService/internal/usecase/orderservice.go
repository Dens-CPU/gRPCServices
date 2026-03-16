package usecase

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"

	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

type Storage interface {
	AddOrderStorage(context.Context, order.Order, []order.Market) (string, string, error) //Добавление нового заказа в хранилище
	GetOrderState(context.Context, order.Key) (string, error)                             //Получение статуса заказа
	ControlOrder(string, int64, string) chan string
}

type MarketsService interface {
	GetEnableMarkets(context.Context) ([]order.Market, error) //Получение списка доступных рынков
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
	logger *zap.Logger
	trace  *trace.TracerProvider
}

func NewOrderServ(in_memory Storage, markets_service MarketsService, notify Notify, logger *zap.Logger, trace *trace.TracerProvider) *OrderService {
	return &OrderService{in_memory, markets_service, notify, logger, trace}
}
