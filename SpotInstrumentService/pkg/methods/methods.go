// В данном файле прописаны методы используемые сервисом SpotInstrumentService
package spotmethods

import (
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	spotmemory "Academy/gRPCServices/SpotInstrumentService/pkg/memory"
	"context"
	"errors"
)

// Интерфейс для рынков
type MarketRepository interface {
	GetEnableMarkets() []*spotmemory.Market //Получение доступных рынков
}

type SpotInstrument struct {
	spotAPI.UnimplementedSpotInstrumentServiceServer
	repo MarketRepository //Абстракия над хранилищами. Не важно, что это будет за хранилище, важно, чтобы оно удовлетворяло всем методам интерфейса.
}

// Конструктор для SpotInstrument
func NewSpotInstrument(repo MarketRepository) *SpotInstrument {
	return &SpotInstrument{repo: repo}
}

// Список доступных рынков
func (s *SpotInstrument) ViewMarket(ctx context.Context, req *spotAPI.ViewReq) (*spotAPI.ViewResp, error) {
	enableMarkets := s.repo.GetEnableMarkets()
	if len(enableMarkets) != 0 {
		return Mapper(enableMarkets), nil
	}
	return &spotAPI.ViewResp{}, errors.New("нет доступных рынков")
}

// Маппер для ViewMarket
func Mapper(em []*spotmemory.Market) *spotAPI.ViewResp {
	var resp spotAPI.ViewResp
	for _, el := range em {
		resp.EnableMarkets = append(resp.EnableMarkets, el.ID)
	}
	return &resp
}
