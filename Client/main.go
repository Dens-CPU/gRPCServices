package main

import (
	"io"

	"context"
	"fmt"
	"log"
	"time"

	"github.com/DencCPU/gRPCServices/OrderService/pkg/orderclient"
	order "github.com/DencCPU/gRPCServices/Protobuf/gen/order_service"
	spot "github.com/DencCPU/gRPCServices/Protobuf/gen/spot_service"
	opentelemetry "github.com/DencCPU/gRPCServices/Shared/opentelimetry"
	"github.com/DencCPU/gRPCServices/SpotInstrumentService/pkg/spotclient"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/metadata"
)

func main() {

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{}))

	tp, err := opentelemetry.NewTrace(context.Background(), "Client", "localhost", "4317")
	if err != nil {
		log.Fatal(err)
	}
	defer tp.Shutdown(context.Background())

	metric, _ := opentelemetry.NewMetricPrometeus(context.Background(), "Client")
	defer metric.Shutdown(context.Background())

	orderclient, err := orderclient.NewClient()
	if err != nil {
		log.Fatal("Инициализация клиента:", err)
	}

	spotclient, err := spotclient.NewClient()
	if err != nil {
		log.Fatal("Ошибка инициализации spot_client:", err)
	}
	spotresp, err := spotclient.ViewMarket(context.Background(), &spot.ViewReq{})
	if err != nil {
		log.Fatal("Ошибка получения ответа от View markets:", err)
	}
	var (
		userID   string
		marketID string
		price    string
		quantity int64
	)
	enableMarkets := spotresp.EnableMarkets
	marketID = enableMarkets[0].MarketId
	fmt.Println("Укажите price:")
	fmt.Scanln(&price)
	fmt.Println("Укажите кол-во товара:")
	fmt.Scanln(&quantity)

	tr := otel.Tracer("main-client")

	ctx, span := tr.Start(context.Background(), "CreateOrder")
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	userID = uuid.New().String()
	md := metadata.Pairs(
		"Test", "Hello",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	createResp, err := orderclient.CreateOrder(ctx, &order.CreateOrderReq{UserId: userID, MarketId: marketID, OrderType: order.OrderType_ORDER_TYPE_EXPRESS, Price: price, Quantity: quantity})
	if err != nil {
		log.Fatal(err)
	}

	orderId := createResp.OrderId
	fmt.Printf("Заказ:%s. Статус:%s.\n", orderId, createResp.OrderStatus)

	ctx, span = tr.Start(context.Background(), "StreamStatus")
	defer span.End()
	streamResp, err := orderclient.StreamOrderUpdate(ctx, &order.StreamOrderUpdateReq{OrderId: orderId, UserId: userID})
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
	time.Sleep(time.Second)
	ctx, span = tr.Start(context.Background(), "GetOrder")
	defer span.End()
	getResp, err := orderclient.GetOrderStatus(ctx, &order.GetOrderReq{OrderId: orderId, UserId: userID})
	if err != nil {
		log.Fatal("Ошибка при запросе GetOrderStatus:", err)
	}
	fmt.Println("Статус заказа:", getResp.OrderStatus)
}
