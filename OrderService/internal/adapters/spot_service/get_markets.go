package spotservice

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	"context"
)

func (s *Client) GetEnableMarkets(ctx context.Context) ([]order.Market, error) {
	resp, err := s.ViewMarket(ctx, &spotAPI.ViewReq{})
	if err != nil {
		return nil, err
	}

	var output = make([]order.Market, 0, len(resp.EnableMarkets))
	for _, em := range resp.EnableMarkets {
		market := order.Market{ID: em.MarketId, Name: em.MarketName}
		output = append(output, market)
	}

	return output, nil
}
