package orderhandlers

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	orderAPI "Academy/gRPCServices/Protobuf/gen/order"
	"context"
)

func (h *Handlers) CreateOrder(ctx context.Context, req *orderAPI.CreateReq) (*orderAPI.CreateResp, error) {
	//Добавить валидацию запроса
	newOrder := order.Order{
		User_id:    req.UserId,
		Market_id:  req.MarketId,
		Order_type: req.OrderType,
		Price:      req.Price,
		Quantity:   req.Quantity,
	}
	orderID, status, err := h.orderService.CreateOrder(ctx, newOrder)
	if err != nil {
		return &orderAPI.CreateResp{}, err
	}
	resp := orderAPI.CreateResp{OrderId: orderID, Status: status}
	return &resp, nil
}
