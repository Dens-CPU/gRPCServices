package orderhandlers

import (
	"context"
	"fmt"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	order "github.com/DencCPU/gRPCServices/Protobuf/gen/order_service"
)

func (h *Handlers) GetOrderStatus(ctx context.Context, req *order.GetOrderReq) (*order.GetOrderResp, error) {
	//Валидация запроса
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("uncorrect request format:%w", err)
	}
	//Добавить валидацию запроса
	key := orderdomain.Key{
		OrderId: req.OrderId,
		UserId:  req.UserId,
	}
	status, err := h.Service.GetStatus(ctx, key)
	if err != nil {
		return &order.GetOrderResp{}, err
	}
	resp := order.GetOrderResp{OrderStatus: status, OrderId: key.OrderId}
	return &resp, nil
}
