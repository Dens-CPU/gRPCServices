package userhandlers

import (
	"context"
	"errors"
	"fmt"

	user "github.com/DencCPU/gRPCServices/Protobuf/gen/user_service"
	sharederrors "github.com/DencCPU/gRPCServices/Shared/errors"
	tokensdto "github.com/DencCPU/gRPCServices/UserService/internal/adapters/dto/tokens"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) UpdateTokens(ctx context.Context, req *user.UpdateTokensReq) (*user.UpdateTokensResp, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("incorrect data format:%w", err)
	}

	input := tokensdto.NewInputTokens(req.AccessToken, req.RefreshToken)
	pairToken, err := h.Service.UpdateTokens(ctx, input)
	if err != nil {

		if errors.Is(err, sharederrors.ReAutentification) {
			return nil, status.Error(codes.Unauthenticated, sharederrors.ReAutentification.Error())
		}

		return &user.UpdateTokensResp{}, err
	}

	resp := &user.UpdateTokensResp{
		AccessToken:  pairToken.AccessToken,
		RefreshToken: pairToken.RefreshToken,
		ExpireAt:     timestamppb.New(pairToken.ExpireAt),
	}
	return resp, nil
}
