package usecase

import (
	"context"
	"fmt"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
)

// Получение статуса заказа в стриминге
func (o *OrderService) StreamGetState(ctx context.Context, key orderdomain.Key) chan string {

	//Добавление новой подписки для получения статусов
	stateCh := o.Notify.AddNewSub(key)
	fmt.Println("Подписан новый клиент")

	//Получение кол-ва каналов
	quantiryCh := o.GetNumbersSubsChan(key)
	if quantiryCh == 1 {
		o.UpdateStatusSubs(key)
		fmt.Println("Запущена служба")
	}

	return stateCh
}
