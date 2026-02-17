package main

import (
	"Academy/gRPCServices/OrderService/internal/adapters/memory"
	spotservice "Academy/gRPCServices/OrderService/internal/adapters/spot_service"
	orderhandlers "Academy/gRPCServices/OrderService/internal/controllers/grpc_handlers"
	"Academy/gRPCServices/OrderService/internal/usecase"
	"Academy/gRPCServices/OrderService/pkg/orderserver"
	orderAPI "Academy/gRPCServices/Protobuf/gen/order"
	"fmt"
	"log"
)

func main() {
	storage := memory.NewStorage() //Инициализация хранилища in-memory

	// ctx := context.Background()
	// storage, err := postgres.NewDB(ctx) //Инициализаця хранилища postgres
	// if err != nil {
	// 	log.Fatal(err)
	// }

	spotClient, err := spotservice.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	service := usecase.NewOrderServ(storage, spotClient)

	handlers := orderhandlers.NewHandlers(service)

	grpcServer, err := orderserver.New()
	if err != nil {
		log.Fatal(err)
	}

	orderAPI.RegisterOrderServiceServer(grpcServer, handlers)

	fmt.Println("Сервер работает на порту 8081...")
	err = grpcServer.Serve(grpcServer.Listener)
	if err != nil {
		log.Fatal("Ошибка работы сервера:", err)
	}
}
