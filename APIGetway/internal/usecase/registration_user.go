package usecase

import (
	"context"

	"github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/jwt"
	userdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/user"
	"go.uber.org/zap"
)

func (s *Service) RegistrationUser(ctx context.Context, newUser userdomain.User) (jwt.PairToken, error) {
	pairToken, err := s.client.RegistrationUser(ctx, newUser)
	if err != nil {
		s.logger.Error("ошибка регистрации нового пользователя на user-сервисе:",
			zap.Error(err),
		)
		return jwt.PairToken{}, err
	}
	return pairToken, nil
}
