package spothandlers

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	domainusers "Academy/gRPCServices/SpotInstrumentService/internal/domain/users"
	"Academy/gRPCServices/SpotInstrumentService/internal/usecase"
	"context"
)

type Service interface {
	ViewMarket(context.Context, *domainusers.User) ([]int64, error)
}

type Handlers struct {
	spotAPI.UnimplementedSpotInstrumentServiceServer
	Service Service //Функционал обработчиков
}

// Конструктор для SpotInstrument
func NewHandlers(spotService *usecase.SpotService) *Handlers {
	return &Handlers{Service: spotService}
}
