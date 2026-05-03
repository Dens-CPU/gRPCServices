package userclient

import (
	user "github.com/DencCPU/gRPCServices/Protobuf/gen/user_service"
	"github.com/DencCPU/gRPCServices/UserService/pkg/userclient"
	"github.com/sony/gobreaker"
)

type Client struct {
	user.UserServiceClient
	breaker *gobreaker.CircuitBreaker
}

func NewClient(breaker *gobreaker.CircuitBreaker) (*Client, error) {
	client, err := userclient.NewClient()
	if err != nil {
		return nil, err
	}
	return &Client{client, breaker}, nil
}
