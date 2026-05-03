package spothandlers

import (
	"context"
	"fmt"

	spot "github.com/DencCPU/gRPCServices/Protobuf/gen/spot_service"
	domainusers "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/domain/users"
)

func (h *Handlers) ViewMarket(ctx context.Context, req *spot.ViewReq) (*spot.ViewResp, error) {
	//Validation
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid request format:%w", err)
	}

	input := domainusers.Input{
		UserRole:  domainusers.UserRole(req.UserRoles),
		UserId:    req.UserId,
		PageSize:  int(req.PageSize),
		PageToken: req.PageToken,
	}

	output, pageToken, err := h.Service.ViewMarket(ctx, input)
	if err != nil {
		return &spot.ViewResp{}, err
	}

	//Формирование ответа
	resp := &spot.ViewResp{}
	resp.EnableMarkets = make([]*spot.Markets, 0, len(output))

	for _, el := range output {
		market := spot.Markets{MarketId: el.ID, MarketName: el.Name}
		resp.EnableMarkets = append(resp.EnableMarkets, &market)
	}
	resp.PageToken = pageToken
	return resp, nil
}
