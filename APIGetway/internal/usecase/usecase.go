package usecase

import (
	"context"

	"github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/jwt"
	userdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/user"
	"go.uber.org/zap"
)

type UserClient interface {
	RegistrationUser(ctx context.Context, newUser userdomain.User) (jwt.PairToken, error)
}
type Service struct {
	client UserClient
	logger *zap.Logger
}

func NewService(client UserClient, logger *zap.Logger) *Service {
	return &Service{client, logger}
}
