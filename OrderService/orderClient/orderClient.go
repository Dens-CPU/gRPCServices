package main

import (
	orderAPI "Academy/gRPCServices/Protobuf/order"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Ошибка подклюючния к сервису OrderServ")
	}

	ctx := context.Background()

	orderClient := orderAPI.NewOrderServiceClient(conn) //Контекст. Не работат WithTimeOut?

	var (
		userID    int64
		marketID  int64
		orderType string
		price     float64
		quantity  int64
	)

	fmt.Println("Доступные рынки:\n1.Yandex Market;\n2.OZON;\n3.Wildberris;\n4.Aliexpress.")
	fmt.Print("Введите user_id:")
	fmt.Scanln(&userID)
	fmt.Println("Введите market_id из предоставленного списка:")
	fmt.Scanln(&marketID)
	fmt.Println("Укажите тип заказа (normal,express):")
	fmt.Scanln(&orderType)
	fmt.Println("Укажите price:")
	fmt.Scanln(&price)
	fmt.Println("Укажите кол-во товара:")
	fmt.Scanln(&quantity)

	resp, err := orderClient.CreateOrder(ctx, &orderAPI.CreateReq{UserId: userID, MarketId: marketID, OrderType: orderType, Price: price, Quantity: quantity})
	if err != nil {
		log.Fatal("Ошибка при запросе CreateOrder:", err)
	}

	orderId := resp.OrderId
	fmt.Printf("Заказ:%d. Статус:%s.\n", orderId, resp.Status)

	time.Sleep(3 * time.Second)
	getResp, err := orderClient.GetOrderStatus(ctx, &orderAPI.GetReq{OrderId: orderId, UserId: userID})
	if err != nil {
		log.Fatal("Ошибка при запросе GetOrderStatus:", err)
	}
	fmt.Println("Статус заказа:", getResp.OrderStatus)
}
