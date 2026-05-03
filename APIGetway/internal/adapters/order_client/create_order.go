package orderclient

import (
	"context"
	"errors"
	"fmt"

	orderdto "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/order_service"
	orderdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/order"
	"github.com/DencCPU/gRPCServices/Protobuf/gen/common"
	"github.com/DencCPU/gRPCServices/Protobuf/gen/order_service"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) CreateNewOrder(ctx context.Context, order orderdomain.OrderInfo) (orderdto.Output, error) {
	result, err := c.breaker.Execute(func() (interface{}, error) {
		req := order_service.CreateOrderReq{
			UserId:   order.UserId,
			MarketId: order.MarketId,
			Price:    order.Price,
			Quantity: order.Quantity,
			UserRole: common.UserRole(order.UserRole),
		}
		switch order.Order_type {
		case "normal":
			req.OrderType = order_service.OrderType_ORDER_TYPE_NORMAL
		case "express":
			req.OrderType = order_service.OrderType_ORDER_TYPE_EXPRESS
		default:
			return orderdto.Output{}, errors.New("incorrect oreder type")
		}

		resp, err := c.CreateOrder(ctx, &req)
		if err != nil {
			return orderdto.Output{}, err
		}

		if resp.OrderId == "" || resp.OrderStatus == "" {
			return orderdto.Output{}, errors.New("incorrect response from the server")
		}
		return resp, nil
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return orderdto.Output{}, status.Errorf(codes.Unavailable, "service is temporarily unavailable")
		}
		return orderdto.Output{}, err
	}

	resp, ok := result.(*order_service.CreateOrderResp)
	if !ok {
		return orderdto.Output{}, fmt.Errorf("Inappropriate result type:%T", result)
	}

	output := orderdto.NewOutput(resp.OrderId, resp.OrderStatus)
	return output, nil
}
