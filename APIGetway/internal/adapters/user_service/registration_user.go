package userservice

import (
	"context"

	"github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/dto/jwt"
	userdomain "github.com/DencCPU/gRPCServices/APIGetway/internal/domain/user"
	"github.com/DencCPU/gRPCServices/Protobuf/gen/user_service"
)

func (c *Client) RegistrationUser(ctx context.Context, newUser userdomain.User) (jwt.PairToken, error) {
	req := &user_service.CreateUserReq{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserRole: user_service.UserRole_USER_ROLE_BASIC_USER,
	}
	resp, err := c.CreateUser(ctx, req)
	if err != nil {
		return jwt.PairToken{}, err
	}
	pairToken := jwt.NewPairToken(
		resp.AccsessToken,
		resp.RefreshToken,
		resp.ExpireAt.AsDuration(),
	)
	return pairToken, nil
}
