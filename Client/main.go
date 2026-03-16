package main

import (
	"Academy/gRPCServices/OrderService/pkg/orderclient"
	orderAPI "Academy/gRPCServices/Protobuf/gen/order"
	"Academy/gRPCServices/Shared/opentelimetry"

	"io"

	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{}))

	tp, err := opentelimetry.NewTrace(context.Background(), "Client")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	metric, _ := opentelimetry.NewMetricPrometeus(context.Background(), "Client")
	defer metric.Shutdown(context.Background())

	client, err := orderclient.NewClient()
	if err != nil {
		log.Fatal("Инициализация клиента:", err)
	}
	var (
		userID    int64
		marketID  int64
		orderType string
		price     string
		quantity  int64
	)

	fmt.Println("Доступные рынки:\n0.Yandex Market;\n1.OZON;\n2.Wildberris;\n3.Aliexpress.")
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

	tr := otel.Tracer("main-client")

	ctx, span := tr.Start(context.Background(), "CreateOrder")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	createResp, err := client.CreateOrder(ctx, &orderAPI.CreateReq{UserId: userID, MarketId: marketID, OrderType: orderType, Price: price, Quantity: quantity})
	if err != nil {
		log.Fatal(err)
	}

	orderId := createResp.OrderId
	fmt.Printf("Заказ:%s. Статус:%s.\n", orderId, createResp.Status)

	ctx, span = tr.Start(context.Background(), "StreamStatus")
	defer span.End()
	streamResp, err := client.StreamOrderUpdate(ctx, &orderAPI.GetReq{OrderId: orderId, UserId: userID})
	for {
		resp, err := streamResp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.OrderStatus)
	}

	time.Sleep(3 * time.Second)
	ctx, span = tr.Start(context.Background(), "GetOrder")
	defer span.End()
	getResp, err := client.GetOrderStatus(ctx, &orderAPI.GetReq{OrderId: orderId, UserId: userID})
	if err != nil {
		log.Fatal("Ошибка при запросе GetOrderStatus:", err)
	}
	fmt.Println("Статус заказа:", getResp.OrderStatus)
}
