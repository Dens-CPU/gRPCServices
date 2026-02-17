package spothandlers

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot_service"
	"Academy/gRPCServices/SpotInstrumentService/internal/usecase"
)

type Handlers struct {
	spotAPI.UnimplementedSpotInstrumentServiceServer
	spotService *usecase.SpotService
}

// Конструктор для SpotInstrument
func NewHandlers(spotService *usecase.SpotService) *Handlers {
	return &Handlers{spotService: spotService}
}
