package main

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot_service"
	"Academy/gRPCServices/SpotInstrumentService/internal/adapters/memory"
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
	spothandlers "Academy/gRPCServices/SpotInstrumentService/internal/controllers/grpc_handlers"
	grpcserver "Academy/gRPCServices/SpotInstrumentService/pkg/spotserver"
	"context"
	"time"

	"Academy/gRPCServices/SpotInstrumentService/internal/usecase"

	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//Инициализация хранилища
	storage := memory.NewStorage()
	err := storage.AddMarkets()
	if err != nil {
		log.Fatal("ошибка добавления рынков:", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		l := storage.AccessControl() //Запуск управления рынками
		log.Println("управление рынками:", l)
	}()

	service := usecase.NewSpotInstrument(storage)
	handlers := spothandlers.NewHandlers(service)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	redis, err := redisadapter.NewRedis(ctx)

	grpcServer, err := grpcserver.New(redis)

	spotAPI.RegisterSpotInstrumentServiceServer(grpcServer, handlers)

	fmt.Println("Сервер запущен на порту 8080...")
	err = grpcServer.Serve(grpcServer.Listener)
	if err != nil {
		log.Fatal("ошибка работы сервера:", err)
	}
	wg.Wait()
}
