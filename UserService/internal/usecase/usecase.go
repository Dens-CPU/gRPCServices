package usecase

import (
	"context"
	"time"

	domainuser "github.com/DencCPU/gRPCServices/UserService/internal/domain/user"
	"go.uber.org/zap"
)

type Storage interface {
	AddUser(context.Context, domainuser.User) (string, string, error)
	UpdatePassword(context.Context, string, string) error
}
type JWT interface {
	CreateAccsesToken(string, string) (string, time.Duration, error)
}

type Service struct {
	Storage
	JWT
	logger *zap.Logger
}

func NewService(storage Storage, logger *zap.Logger, jwt JWT) *Service {
	return &Service{storage, jwt, logger}
}
