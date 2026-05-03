package orderhandlers

import (
	"context"
	"fmt"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	order "github.com/DencCPU/gRPCServices/Protobuf/gen/order_service"

	"github.com/shopspring/decimal"
)

func (h *Handlers) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	//Validation
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request format:%w", err)
	}

	//Parse string to decimal
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return nil, fmt.Errorf("invalid price format:%w", err)
	}

	newOrder := orderdomain.Order{
		UserId:    req.UserId,
		MarketId:  req.MarketId,
		OrderType: orderdomain.OrderType(req.OrderType),
		Price:     price,
		Quantity:  req.Quantity,
		UserRole:  orderdomain.UserRole(req.UserRole),
	}

	if newOrder.OrderType == orderdomain.ORDER_TYPE_UNSPECIFIED {
		return nil, fmt.Errorf("unknow order type")
	}
	orderID, status, err := h.Service.CreateOrder(ctx, newOrder)
	if err != nil {
		return &order.CreateOrderResp{}, err
	}

	resp := order.CreateOrderResp{OrderId: orderID, OrderStatus: status}
	if err = resp.Validate(); err != nil {
		return nil, fmt.Errorf("invalid response format:%w", err)
	}
	return &resp, nil
}
