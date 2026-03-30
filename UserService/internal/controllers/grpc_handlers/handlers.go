package userhandlers

import (
	"context"
	"time"

	"github.com/DencCPU/gRPCServices/Protobuf/gen/user_service"
	domainuser "github.com/DencCPU/gRPCServices/UserService/internal/domain/user"
)

type Service interface {
	CreateUser(context.Context, domainuser.User) (string, string, time.Duration, error)
	UpdatePassword(context.Context, string, string) error
}

type Handler struct {
	user_service.UnimplementedUserServiceServer
	Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}
