package main

import (
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Ошибка подключения к серверу:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := spotAPI.NewSpotInstrumentServiceClient(conn)

	resp, err := client.ViewMarket(ctx, &spotAPI.ViewReq{})
	if err != nil {
		fmt.Println("Ошибка при получение ответа от сервреа:", err)
	}
	enableMarkets := resp.EnableMarkets
	for _, m := range enableMarkets {
		fmt.Println(m)
	}
}
