package usecase

import (
	"context"

	"github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/tokens"
)

func (s *Service) Authentication(ctx context.Context, email, password string) (tokens.PairToken, error) {
	ctx, span := s.tracer.Start(ctx, "Autentication")
	defer span.End()

	pairToken, err := s.user_client.AuthenticationUser(ctx, email, password)
	if err != nil {
		return tokens.PairToken{}, err
	}
	return pairToken, nil
}
