package spotmethods_test

import (
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	spotmemory "Academy/gRPCServices/SpotInstrumentService/pkg/memory"
	spotmethods "Academy/gRPCServices/SpotInstrumentService/pkg/methods"
	"context"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestSpotInstrument_ViewMarket(t *testing.T) {
	markets := []string{"Yandex Market", "Ozon", "Wilderris", "Aliexpress"}
	storage := spotmemory.NewStorage(len(markets))
	go func() {
		storage.AccessControl(markets)
	}()
	service := spotmethods.NewSpotInstrument(storage)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Error("ошибка ннастройки сервера:", err)
	}
	grpcServer := grpc.NewServer()
	spotAPI.RegisterSpotInstrumentServiceServer(grpcServer, service)
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			t.Error("ошибка работы сервера:", err)
			return
		}
	}()
	conn, err := grpc.NewClient(":8080", grpc.WithInsecure())
	if err != nil {
		t.Error("ошибка подключения к серверу:", err)
	}
	ctx := context.Background

	time.Sleep(time.Second)
	client := spotAPI.NewSpotInstrumentServiceClient(conn)
	resp, err := client.ViewMarket(ctx(), &spotAPI.ViewReq{})
	if err != nil {
		t.Error("Ошибка запроса ViewMarket", err)
		return
	}
	em := resp.EnableMarkets
	if len(em) == 0 {
		t.Errorf("ошибка: ожидалась длина не 0")
	}
}
