package usecase

import (
	domainerrors "Academy/gRPCServices/SpotInstrumentService/internal/domain/errors"
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
	domainusers "Academy/gRPCServices/SpotInstrumentService/internal/domain/users"
)

// Получение доступных рынков
func (s *SpotService) ViewMarket(user *domainusers.User) ([]int64, error) {
	var enableMarkets []*domainmarket.Market
	enableMarkets = s.GetEnableMarkets()

	if len(enableMarkets) == 0 {
		return nil, domainerrors.Avalible_markets
	}

	return Mapper(enableMarkets), nil
}

// Маппер для ViewMarket
func Mapper(em []*domainmarket.Market) []int64 {
	var resp []int64
	for _, el := range em {
		resp = append(resp, el.ID)
	}
	return resp
}
