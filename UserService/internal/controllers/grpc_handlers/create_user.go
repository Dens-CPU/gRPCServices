package userhandlers

import (
	"context"
	"errors"
	"fmt"

	user "github.com/DencCPU/gRPCServices/Protobuf/gen/user_service"
	jwtdto "github.com/DencCPU/gRPCServices/UserService/internal/adapters/dto/tokens"
	domainuser "github.com/DencCPU/gRPCServices/UserService/internal/domain/user"
	"google.golang.org/protobuf/types/known/durationpb"
)

func (h *Handler) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("неправильный формат данных:%w", err)
	}
	var role string

	switch req.UserRole {
	case user.UserRole_USER_ROLE_BASIC_USER:
		role = "basic"
	case user.UserRole_USER_ROLE_PREMIUM_USER:
		role = "premium"
	default:
		return nil, errors.New("типа пользователя не существует")
	}

	newUser := domainuser.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     role,
	}
	refreshToken, accsessToken, TTL, err := h.Service.CreateUser(ctx, newUser)
	pairToken := jwtdto.NewPairToken(accsessToken, refreshToken, TTL)

	resp := &user.CreateUserResp{
		AccsessToken: pairToken.AccsesToken,
		RefreshToken: pairToken.RefreshToken,
		ExpireAt:     durationpb.New(pairToken.Expire_at),
	}
	return resp, nil
}
