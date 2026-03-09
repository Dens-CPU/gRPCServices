package usecase

import (
	domainerrors "Academy/gRPCServices/SpotInstrumentService/internal/domain/errors"
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
	domainusers "Academy/gRPCServices/SpotInstrumentService/internal/domain/users"
	"context"

	"go.uber.org/zap"
)

// Получение доступных рынков
func (s *SpotService) ViewMarket(ctx context.Context, user *domainusers.User) ([]int64, error) {

	var enableMarkets []*domainmarket.Market
	tracer := s.trace.Tracer("SpotSevrvice")
	ctx, span := tracer.Start(ctx, "get enable markets")
	defer span.End()

	enableMarkets = s.GetEnableMarkets()

	if len(enableMarkets) == 0 {
		s.logger.Error("Нет доступных рынков")
		return nil, domainerrors.Avalible_markets
	}
	s.logger.Info("Список доступных рынков получен",
		zap.String("get enable markets span:", span.SpanContext().TraceID().String()),
	)
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
