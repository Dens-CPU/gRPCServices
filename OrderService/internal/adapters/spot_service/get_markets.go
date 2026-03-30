package spotservice

import (
	"context"
	"errors"
	"fmt"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	spot "github.com/DencCPU/gRPCServices/Protobuf/gen/spot_service"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Client) GetEnableMarkets(ctx context.Context) ([]orderdomain.Market, error) {
	//Создание оболочки брейкера над зопросом
	result, err := s.breaker.Execute(func() (interface{}, error) {

		//Запрос к сервису
		resp, err := s.ViewMarket(ctx, &spot.ViewReq{})
		if err != nil {
			return nil, err
		}
		return resp, nil
	},
	)

	//Проверка состояния брейкера (закрыт, открыт)
	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, status.Errorf(codes.Unavailable, "сервис временно недоступен")
		}
		return nil, err
	}

	//Приведение результата безопасного запроса к определенному типу
	resp, ok := result.(*spot.ViewResp)
	if !ok {
		return nil, fmt.Errorf("Несоответсвующий тип result:%T", result)
	}

	//Проверка ответа от сервера
	if len(resp.EnableMarkets) == 0 {
		return []orderdomain.Market{}, nil
	}

	//Формирование ответа
	var output = make([]orderdomain.Market, 0, len(resp.EnableMarkets))
	for _, em := range resp.EnableMarkets {
		market := orderdomain.Market{ID: em.MarketId, Name: em.MarketName}
		output = append(output, market)
	}
	return output, nil
}
