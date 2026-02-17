package spothandlers

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot_service"
	domainusers "Academy/gRPCServices/SpotInstrumentService/internal/domain/users"
	"context"
)

func (h *Handlers) ViewMarket(ctx context.Context, req *spotAPI.ViewReq) (*spotAPI.ViewResp, error) {
	user := domainusers.NewUser(domainusers.UserType(req.UserRoles))
	output, err := h.spotService.ViewMarket(user)
	if err != nil {
		return &spotAPI.ViewResp{}, err
	}
	resp := &spotAPI.ViewResp{EnableMarkets: output}
	return resp, nil
}
