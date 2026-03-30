package usecase

import (
	"context"
	"time"

	domainuser "github.com/DencCPU/gRPCServices/UserService/internal/domain/user"
	"go.uber.org/zap"
)

func (s *Service) CreateUser(ctx context.Context, user domainuser.User) (string, string, time.Duration, error) {

	//Добавление пользователя в БД
	user_id, refreshToken, err := s.AddUser(ctx, user)
	if err != nil {
		s.logger.Error("ошибка добавления в пользователя БД:",
			zap.Error(err),
		)
		return "", "", 0, err
	}
	accsessToken, TTL, err := s.CreateAccsesToken(user_id, user.Email)
	if err != nil {
		s.logger.Error("Ошибка формирования jwt:",
			zap.Error(err),
		)
		return "", "", 0, err
	}
	return refreshToken, accsessToken, TTL, nil
}
