package spotservice

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot_service"
	"Academy/gRPCServices/SpotInstrumentService/pkg/spotclient"
)

type Client struct {
	spotAPI.SpotInstrumentServiceClient
}

func NewClient() (*Client, error) {
	client, err := spotclient.NewClient()
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
