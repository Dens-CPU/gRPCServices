package orderhandlers

import (
	"context"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	"github.com/DencCPU/gRPCServices/OrderService/internal/usecase"
	order "github.com/DencCPU/gRPCServices/Protobuf/gen/order_service"
)

type Service interface {
	CreateOrder(context.Context, orderdomain.Order) (string, string, error)
	GetStatus(context.Context, orderdomain.Key) (string, error)
	StreamGetState(ctx context.Context, key orderdomain.Key) chan string
}

type Handlers struct {
	order.UnimplementedOrderServiceServer
	Service Service
}

func NewHandlers(orderService *usecase.OrderService) *Handlers {
	return &Handlers{Service: orderService}
}
