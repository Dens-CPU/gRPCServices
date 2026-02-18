package usecase

import (
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
)

// Интерфейс для рынков
type StorageRepo interface {
	GetEnableMarkets() []*domainmarket.Market //Получение доступных рынков
}

type SpotService struct {
	StorageRepo //Абстракия над хранилищами. Не важно, что это будет за хранилище, важно, чтобы оно удовлетворяло всем методам интерфейса.
}

// Конструктор для SpotInstrument
func NewSpotInstrument(repo StorageRepo) *SpotService {
	return &SpotService{repo}
}
