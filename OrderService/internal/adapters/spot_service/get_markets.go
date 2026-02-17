package spotservice

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot_service"
	"context"
)

func (s *Client) GetEnableMarkets(ctx context.Context) ([]int64, error) {
	resp, err := s.ViewMarket(ctx, &spotAPI.ViewReq{})
	if err != nil {
		return nil, err
	}
	return resp.EnableMarkets, nil
}
