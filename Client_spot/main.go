package main

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	"Academy/gRPCServices/Shared/opentelimetry"
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
)

func main() {

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{}))

	trace, err := opentelimetry.NewTrace(context.Background(), "Client_spot")
	if err != nil {
		log.Fatal(err)
	}

	defer trace.Shutdown(context.Background())
	// Контекст с таймаутом для подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Подключение к серверу gRPC с OTel stats handler
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithInsecure(),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()), // <-- OTel клиентский handler
	)
	if err != nil {
		log.Fatal("ошибка подключения:", err)
	}
	defer conn.Close()

	// Создаём клиент
	client := spotAPI.NewSpotInstrumentServiceClient(conn)
	tracer := trace.Tracer("Client_spot")
	ctx, span := tracer.Start(ctx, "View market")
	defer span.End()

	// Делаем запрос
	resp, err := client.ViewMarket(ctx, &spotAPI.ViewReq{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Доступные рынки:", resp.EnableMarkets)
}
