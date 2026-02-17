package spotclient

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot_service"
	serverconfig "Academy/gRPCServices/SpotInstrumentService/config/server"

	"google.golang.org/grpc"
)

type Client struct {
	spotAPI.SpotInstrumentServiceClient
}

func NewClient() (*Client, error) {
	cfg, err := serverconfig.NewConfig()
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(cfg.Server.Host+cfg.Server.Port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := spotAPI.NewSpotInstrumentServiceClient(conn)
	return &Client{client}, nil
}
