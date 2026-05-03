package usecase

import (
	"context"

	tokensdto "github.com/DencCPU/gRPCServices/UserService/internal/adapters/dto/tokens"
	"go.uber.org/zap"
)

func (s *Service) AuthenticationUser(ctx context.Context, email, password string) (tokensdto.PairToken, error) {
	ctx, span := s.tracer.Start(ctx, "authentication user")
	defer span.End()

	authUser, err := s.Authentication(ctx, email, password)
	if err != nil {
		s.logger.Error("user authentication error:",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return tokensdto.PairToken{}, err
	}
	s.logger.Info("user confirmed")

	ctx, span = s.tracer.Start(ctx, "update expireAt in refresh token")
	defer span.End()

	refreshToken, err := s.UpdateExpireAt(ctx, authUser.ID)
	if err != nil {
		s.logger.Error("error update expireAt in refresh token:",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return tokensdto.PairToken{}, err
	}
	s.logger.Info("update expireAt in refresh token successfully",
		zap.String("spanID:", span.SpanContext().SpanID().String()),
	)

	ctx, span = s.tracer.Start(ctx, "create access token")
	defer span.End()

	accessToken, expireAt, err := s.CreateAccessToken(authUser.ID, email, authUser.Role)
	if err != nil {
		s.logger.Error("access token creation error",
			zap.String("spanID:", span.SpanContext().SpanID().String()),
			zap.Error(err),
		)
		return tokensdto.PairToken{}, err
	}
	s.logger.Info("authentication successful",
		zap.String("spanID:", span.SpanContext().SpanID().String()),
	)

	pairToken := tokensdto.NewPairToken(accessToken, refreshToken, expireAt)
	return pairToken, nil
}
