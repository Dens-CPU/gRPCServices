package spothandlers

import (
	"context"

	spot "github.com/DencCPU/gRPCServices/Protobuf/gen/spot_service"
	viewdto "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/adapters/dto"
	domainusers "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/domain/users"
	"github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/usecase"
)

type Service interface {
	ViewMarket(context.Context, *domainusers.User) ([]viewdto.Output, error)
}

type Handlers struct {
	spot.UnimplementedSpotInstrumentServiceServer
	Service Service //Функционал обработчиков
}

// Конструктор для SpotInstrument
func NewHandlers(spotService *usecase.SpotService) *Handlers {
	return &Handlers{Service: spotService}
}
