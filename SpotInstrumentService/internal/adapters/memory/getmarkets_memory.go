package memory

import domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"

// Получение доступных рынков
func (s *Storage) GetEnableMarkets() []*domainmarket.Market {
	var enableMarkets []*domainmarket.Market

	for _, value := range s.date { //Поиск доступных рынков в хранилище
		s.mu.RLock()
		if value.Delete_at == nil || value.Enable == true {
			enableMarkets = append(enableMarkets, value)
		}
		s.mu.RUnlock()
	}
	return enableMarkets
}
