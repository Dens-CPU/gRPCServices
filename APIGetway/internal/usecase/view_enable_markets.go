package usecase

import (
	"context"

	spotservicedto "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/spot_service"
	"go.uber.org/zap"
)

func (s *Service) ViewEnableMarkets(ctx context.Context, input spotservicedto.Input) ([]spotservicedto.Output, error) {
	ctx, span := s.tracer.Start(ctx, "View markets")
	defer span.End()

	markets, err := s.spot_client.ViewEnableMarkets(ctx, input)
	if err != nil {
		s.logger.Error("error getting list of markets",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info("list of markets received")
	return markets, nil
}
