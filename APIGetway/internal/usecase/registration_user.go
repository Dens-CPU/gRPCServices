package usecase

import (
	"context"

	"github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/tokens"
	userdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/user"
	"go.uber.org/zap"
)

func (s *Service) RegistrationUser(ctx context.Context, newUser userdomain.User) (tokens.PairToken, error) {
	ctx, span := s.tracer.Start(ctx, "Registration user")
	defer span.End()

	pairToken, err := s.user_client.RegistrationUser(ctx, newUser)
	if err != nil {
		s.logger.Error("error registering a new user",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return tokens.PairToken{}, err
	}

	s.logger.Info("User is registred",
		zap.String("spanID:", span.SpanContext().SpanID().String()),
	)

	return pairToken, nil
}
