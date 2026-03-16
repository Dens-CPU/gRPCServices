package memory

import (
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
)

// Добавление маркетов
func (s *Storage) AddMarkets(markets []string) error {

	//Заполнение storage
	for i, m := range markets {
		s.date[m] = &domainmarket.Market{ID: int64(i), Name: m, Enable: true, Delete_at: nil}
	}
	return nil
}
