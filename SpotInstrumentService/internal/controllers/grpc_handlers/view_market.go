package spothandlers

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	domainusers "Academy/gRPCServices/SpotInstrumentService/internal/domain/users"
	"context"
)

func (h *Handlers) ViewMarket(ctx context.Context, req *spotAPI.ViewReq) (*spotAPI.ViewResp, error) {
	user := domainusers.NewUser(domainusers.UserType(req.UserRoles)) //Преобразование запроса в доменную структуру User

	output, err := h.Service.ViewMarket(user) //Запрос сервиса на получение доступных рынков
	if err != nil {
		return &spotAPI.ViewResp{}, err
	}

	resp := &spotAPI.ViewResp{EnableMarkets: output} //Формирование ответа сервера
	return resp, nil
}
