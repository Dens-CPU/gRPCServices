package orderhandlers

import (
	"context"
	"errors"
	"fmt"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	order "github.com/DencCPU/gRPCServices/Protobuf/gen/order_service"

	"github.com/shopspring/decimal"
)

func (h *Handlers) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	//Валидация запроса
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("неправильный формат запроса:%w", err)
	}

	//Парсинг цены
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return nil, fmt.Errorf("неправильный формат цены:%w", err)
	}
	var orderType string

	switch req.OrderType {
	case order.OrderType_ORDER_TYPE_NORMAL:
		orderType = "normal"
	case order.OrderType_ORDER_TYPE_EXPRESS:
		orderType = "express"
	default:
		return &order.CreateOrderResp{}, errors.New("неверный тип заказа")
	}
	//Формирование нового заказа
	newOrder := orderdomain.Order{
		User_id:    req.UserId,
		Market_id:  req.MarketId,
		Order_type: orderType,
		Price:      price,
		Quantity:   req.Quantity,
	}

	//Создание заказа
	orderID, status, err := h.Service.CreateOrder(ctx, newOrder)
	if err != nil {
		return &order.CreateOrderResp{}, err
	}

	//Формирование ответа
	resp := order.CreateOrderResp{OrderId: orderID, OrderStatus: status}
	if err = resp.Validate(); err != nil {
		return nil, fmt.Errorf("неправильный формат ответа сервера:%w", err)
	}
	return &resp, nil
}
