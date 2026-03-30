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
		return nil, fmt.Errorf("неправильный формат запроса:%w", err)
	}
	//Добавить валидацию запроса
	key := orderdomain.Key{
		Order_id: req.OrderId,
		User_id:  req.UserId,
	}
	status, err := h.Service.GetStatus(ctx, key)
	if err != nil {
		return &order.GetOrderResp{}, err
	}
	resp := order.GetOrderResp{OrderStatus: status, OrderId: key.Order_id}
	if err = resp.Validate(); err != nil {
		return nil, fmt.Errorf("неправильный формат ответа сервера:%w", err)
	}
	return &resp, nil
}
