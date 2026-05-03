package orderclient

import (
	"github.com/DencCPU/gRPCServices/OrderService/pkg/orderclient"
	"github.com/DencCPU/gRPCServices/Protobuf/gen/order_service"
	"github.com/sony/gobreaker"
)

type Client struct {
	order_service.OrderServiceClient
	breaker *gobreaker.CircuitBreaker
}

func NewClient(breaker *gobreaker.CircuitBreaker) (*Client, error) {
	client, err := orderclient.NewClient()
	if err != nil {
		return &Client{}, err
	}
	return &Client{client, breaker}, nil
}
