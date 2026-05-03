package usecase

import (
	"context"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
)

type Storage interface {
	AddOrderStorage(ctx context.Context, newOrder orderdomain.Order, markets []orderdomain.Market) (orderID string, orderStatus string, err error) //Добавление нового заказа в хранилище
	GetOrderState(ctx context.Context, key orderdomain.Key) (status string, err error)                                                             //Получение статуса заказа
}

type MarketsService interface {
	GetEnableMarkets(ctx context.Context, userID string, userRole orderdomain.UserRole) ([]orderdomain.Market, error) //Получение списка доступных рынков
}

type Notify interface {
	AddNewState(string, string, chan string)
	GetStatus(orderdomain.Key) string
	AddNewSub(orderdomain.Key) chan string
	GetNumbersSubsChan(orderdomain.Key) int
	UpdateStatusSubs(context.Context, orderdomain.Key)
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
