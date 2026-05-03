package userhandlers

import (
	"context"
	"errors"

	"github.com/DencCPU/gRPCServices/Protobuf/gen/common"
	"github.com/DencCPU/gRPCServices/Protobuf/gen/user_service"
	sharederrors "github.com/DencCPU/gRPCServices/Shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) ValidationTokens(ctx context.Context, req *user_service.ValidationReq) (*user_service.ValidationResp, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	output, err := h.Service.Validation(ctx, req.AccessToken)
	if err != nil {
		if errors.Is(err, sharederrors.ExpiredToken) {
			return nil, status.Error(codes.Unauthenticated, sharederrors.ExpiredToken.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := user_service.ValidationResp{}
	switch output.Role {
	case "basic":
		resp.Role = common.UserRole_USER_ROLE_BASIC_USER
	case "premium":
		resp.Role = common.UserRole_USER_ROLE_PREMIUM_USER

	}
	resp.UserId = output.UserId
	return &resp, nil
}
