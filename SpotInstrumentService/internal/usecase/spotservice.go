package usecase

import (
	domainmarket "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/domain/market"
	domainusers "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/domain/users"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// Интерфейс для рынков
type StorageRepo interface {
	GetEnableMarkets(input domainusers.Input) ([]*domainmarket.Market, string) //Получение доступных рынков
}

type SpotService struct {
	StorageRepo //Абстракия над хранилищами. Не важно, что это будет за хранилище, важно, чтобы оно удовлетворяло всем методам интерфейса.
	logger      *zap.Logger
	tracer      trace.Tracer
}

// Конструктор для SpotInstrument
func NewSpotInstrument(repo StorageRepo, logger *zap.Logger, tracer trace.Tracer) *SpotService {
	return &SpotService{repo, logger, tracer}
}
