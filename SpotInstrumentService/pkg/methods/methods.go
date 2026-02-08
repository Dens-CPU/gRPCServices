// В данном файле прописаны методы используемые сервисом SpotInstrumentService
package spotmethods

import (
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	"Academy/gRPCServices/SpotInstrumentService/pkg/memory"
	"context"
	"errors"
	"sync"
)

type SpotInstrument struct {
	spotAPI.UnimplementedSpotInstrumentServiceServer
	Memory memory.Storage
}

func NewSpotInstrument(size int) *SpotInstrument {
	return &SpotInstrument{Memory: make(memory.Storage)}
}

// Список доступных рынков
func (s *SpotInstrument) ViewMarket(ctx context.Context, req *spotAPI.ViewReq) (*spotAPI.ViewResp, error) {
	var enableMarkets []memory.Market
	var mu sync.RWMutex

	for _, value := range s.Memory {
		mu.RLock()
		if value.Delete_at == nil || value.Enable == true {
			enableMarkets = append(enableMarkets, *value)
		}
		mu.RUnlock()
	}
	if len(enableMarkets) != 0 {
		return Mapper(enableMarkets), nil
	}
	return &spotAPI.ViewResp{}, errors.New("нет доступных рынков")
}

// Маппер для ViewMarket
func Mapper(em []memory.Market) *spotAPI.ViewResp {
	var resp spotAPI.ViewResp
	for _, el := range em {
		resp.EnableMarkets = append(resp.EnableMarkets, el.ID)
	}
	return &resp
}
