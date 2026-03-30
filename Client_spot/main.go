package main

import (
	"context"
	"fmt"
	"log"
	"time"

	spot "github.com/DencCPU/gRPCServices/Protobuf/gen/spot_service"
	opentelemetry "github.com/DencCPU/gRPCServices/Shared/opentelimetry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
)

func main() {

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{}))

	trace, err := opentelemetry.NewTrace(context.Background(), "Client_spot", "localhost", "4317")
	if err != nil {
		log.Fatal(err)
	}

	defer trace.Shutdown(context.Background())
	// Контекст с таймаутом для подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	metrics, err := opentelemetry.NewMetricPrometeus(context.Background(), "SpotClient")
	if err != nil {
		log.Fatal(err)
	}
	defer metrics.Shutdown(context.Background())

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
	client := spot.NewSpotInstrumentServiceClient(conn)
	tracer := trace.Tracer("Client_spot")
	ctx, span := tracer.Start(ctx, "View market")
	defer span.End()

	// Делаем запрос
	resp, err := client.ViewMarket(ctx, &spot.ViewReq{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Доступные рынки:", resp.EnableMarkets)
}
