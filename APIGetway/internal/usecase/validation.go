package usecase

import (
	"context"

	userservicedto "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/user_service"
	"go.uber.org/zap"
)

func (s *Service) Validation(ctx context.Context, accessToken string) (userservicedto.Output, error) {
	ctx, span := s.tracer.Start(ctx, "Validation access token")
	defer span.End()

	output, err := s.user_client.Validation(ctx, accessToken)
	if err != nil {
		s.logger.Error("error validation access token",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return userservicedto.Output{}, err
	}

	s.logger.Info("Validation completed successfully",
		zap.String("spanID:", span.SpanContext().SpanID().String()),
	)

	return output, nil
}
