package spotservice

import (
	orderconfig "github.com/DencCPU/gRPCServices/OrderService/config"
	spot "github.com/DencCPU/gRPCServices/Protobuf/gen/spot_service"
	"github.com/DencCPU/gRPCServices/SpotInstrumentService/pkg/spotclient"
	"github.com/sony/gobreaker"
	"go.uber.org/zap"
)

type Client struct {
	spot.SpotInstrumentServiceClient
	breaker *gobreaker.CircuitBreaker
}

func NewClient(cfg orderconfig.BreakerSetting, logger *zap.Logger, breaker *gobreaker.CircuitBreaker) (*Client, error) {
	client, err := spotclient.NewClient()
	if err != nil {
		return nil, err
	}
	return &Client{client, breaker}, nil
}
