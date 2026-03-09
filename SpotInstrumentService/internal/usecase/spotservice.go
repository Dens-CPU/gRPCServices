package usecase

import (
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"

	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

// Интерфейс для рынков
type StorageRepo interface {
	GetEnableMarkets() []*domainmarket.Market //Получение доступных рынков
}

type SpotService struct {
	StorageRepo //Абстракия над хранилищами. Не важно, что это будет за хранилище, важно, чтобы оно удовлетворяло всем методам интерфейса.
	logger      *zap.Logger
	trace       *trace.TracerProvider
}

// Конструктор для SpotInstrument
func NewSpotInstrument(repo StorageRepo, logger *zap.Logger, trace *trace.TracerProvider) *SpotService {
	return &SpotService{repo, logger, trace}
}
