package spotservice

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	"context"
	"log"
)

func (s *Client) GetEnableMarkets(ctx context.Context) ([]int64, error) {
	log.Println("Сделан запрос в spotService")
	resp, err := s.ViewMarket(ctx, &spotAPI.ViewReq{})
	if err != nil {
		return nil, err
	}
	log.Println("Получен удовлетворительный ответ от сервиса spotService")
	return resp.EnableMarkets, nil
}
