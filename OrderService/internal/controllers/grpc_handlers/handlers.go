package orderhandlers

import (
	"Academy/gRPCServices/OrderService/internal/usecase"
	orderAPI "Academy/gRPCServices/Protobuf/gen/order"
)

type Handlers struct {
	orderAPI.UnimplementedOrderServiceServer
	orderService *usecase.OrderService
}

func NewHandlers(orderService *usecase.OrderService) *Handlers {
	return &Handlers{orderService: orderService}
}
