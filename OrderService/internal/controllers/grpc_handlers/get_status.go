package orderhandlers

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	orderAPI "Academy/gRPCServices/Protobuf/gen/order"
	"context"
)

func (h *Handlers) GetOrderStatus(ctx context.Context, req *orderAPI.GetReq) (*orderAPI.GetResp, error) {
	//Добавить валидацию запроса
	key := order.Key{
		Order_id: req.OrderId,
		User_id:  req.UserId,
	}
	status, err := h.Service.GetStatus(ctx, key)
	if err != nil {
		return &orderAPI.GetResp{}, err
	}
	resp := orderAPI.GetResp{OrderStatus: status}

	return &resp, nil
}
