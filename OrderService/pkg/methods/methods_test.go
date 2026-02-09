package ordermethods_test

import (
	ordermemory "Academy/gRPCServices/OrderService/pkg/memory"
	ordermethods "Academy/gRPCServices/OrderService/pkg/methods"
	orderAPI "Academy/gRPCServices/Protobuf/Order"
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	spotmemory "Academy/gRPCServices/SpotInstrumentService/pkg/memory"
	spotmethods "Academy/gRPCServices/SpotInstrumentService/pkg/methods"
	"context"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestOrderService_CreateOrder(t *testing.T) {
	markets := []string{"Yandex Market", "Ozon", "Wilderris", "Aliexpress"}
	spotStorage := spotmemory.NewStorage(len(markets))
	go func() {
		spotStorage.AccessControl(markets)
	}()
	spotService := spotmethods.NewSpotInstrument(spotStorage)

	spotlis, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Error("ошибка ннастройки сервера:", err)
	}
	spotGrpcServer := grpc.NewServer()
	spotAPI.RegisterSpotInstrumentServiceServer(spotGrpcServer, spotService)
	go func() {
		err := spotGrpcServer.Serve(spotlis)
		if err != nil {
			t.Error("ошибка работы сервера:", err)
			return
		}
	}()

	orderlis, err := net.Listen("tcp", ":8081") //Настройка сервера
	if err != nil {
		t.Error("ошибка настройки сервиса:", err)
		return
	}

	orderGrpcServer := grpc.NewServer()

	orderStorage := ordermemory.NewStorage()
	orderService := ordermethods.NewOrderService(orderStorage)

	orderAPI.RegisterOrderServiceServer(orderGrpcServer, orderService)

	go func() {
		err = orderGrpcServer.Serve(orderlis)
		if err != nil {
			t.Error("Ошибка работы сервера:", err)
			return
		}
	}()

	time.Sleep(time.Second)

	conn, err := grpc.NewClient(":8081", grpc.WithInsecure())
	if err != nil {
		t.Error("ошибка подклюючния к сервису OrderServ:", err)
		return
	}

	ctx := context.Background()

	orderClient := orderAPI.NewOrderServiceClient(conn) //Контекст. Не работат WithTimeOut?
	var userID int64 = 12
	resp, err := orderClient.CreateOrder(ctx, &orderAPI.CreateReq{UserId: userID, MarketId: 1, OrderType: "express", Price: 12.1, Quantity: 5})
	if err != nil {
		t.Error("ошибка выполнения запроса к CreateOrder:", err)
		return
	}
	var orderID = resp.OrderId
	var status = resp.Status
	if orderID == 0 || status == "" {
		t.Errorf("ошибка. ожидалось не %d, и не %s", orderID, status)
		return
	}

	getResp, err := orderClient.GetOrderStatus(ctx, &orderAPI.GetReq{OrderId: orderID, UserId: userID})
	if err != nil {
		t.Error("ошибка выполения запроса к GetOrderStatus:", err)
	}
	status = getResp.OrderStatus
	if status == "" {
		t.Error("ошибка ожидалась не пустая строка")
	}
}
