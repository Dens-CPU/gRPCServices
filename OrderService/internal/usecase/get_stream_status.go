package usecase

import (
	"context"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	"go.uber.org/zap"
)

// Получение статуса заказа в стриминге
func (o *OrderService) StreamGetState(ctx context.Context, key orderdomain.Key) (chan string, error) {

	_, err := o.Storage.GetOrderState(ctx, key)
	if err != nil {
		return nil, err
	}
	//Sign new client
	stateCh := o.Notify.AddNewSub(key)
	o.logger.Info("New client signed",
		zap.String("UserID:", key.UserId),
		zap.String("OrderID:", key.OrderId),
	)

	//Get quantity channels
	quantiryCh := o.GetNumbersSubsChan(key)
	if quantiryCh == 1 {
		o.UpdateStatusSubs(ctx, key)
		o.logger.Info("Service started")
	}

	return stateCh, nil
}
