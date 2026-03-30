package usecase

import (
	"context"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
)

type Storage interface {
	AddOrderStorage(context.Context, orderdomain.Order, []orderdomain.Market) (string, string, error) //Добавление нового заказа в хранилище
	GetOrderState(context.Context, orderdomain.Key) (string, error)                                   //Получение статуса заказа
	ControlOrder(string, string, string) chan string
}

type MarketsService interface {
	GetEnableMarkets(context.Context) ([]orderdomain.Market, error) //Получение списка доступных рынков
}

type Notify interface {
	AddNewState(string, string, chan string)
	GetStatus(orderdomain.Key) string
	AddNewSub(orderdomain.Key) chan string
	GetNumbersSubsChan(orderdomain.Key) int
	UpdateStatusSubs(orderdomain.Key)
}

type OrderService struct {
	Storage
	MarketsService
	Notify
	logger *zap.Logger
	tracer trace.Tracer
}

func NewOrderServ(in_memory Storage, markets_service MarketsService, notify Notify, logger *zap.Logger, trace trace.Tracer) *OrderService {
	return &OrderService{in_memory, markets_service, notify, logger, trace}
}
