package main

import (
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	"Academy/gRPCServices/SpotInstrumentService/pkg/interseptors"
	spotmemory "Academy/gRPCServices/SpotInstrumentService/pkg/memory"
	spotmethods "Academy/gRPCServices/SpotInstrumentService/pkg/methods"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"sync"

	"google.golang.org/grpc"
)

// Добавление маркетов
func AddMarkets(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка прочтения файла:%w", err)
	}
	re := regexp.MustCompile(`([a-zA-Z0-9]+)+\s?([a-zA-Z0-9]+)?`)
	markets := re.FindAllString(string(file), -1)
	return markets, nil
}

func main() {
	path := "./SpotInstrumentService/cmd/markets.txt"
	markets, err := AddMarkets(path)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	//Инициализация хранилища
	storage := spotmemory.NewStorage(len(markets))
	spotService := spotmethods.NewSpotInstrument(storage)
	log.Println("Хранилище инициализированно")

	wg.Add(1)
	go func() {
		defer wg.Done()
		l := storage.AccessControl(markets) //Запуск управления рынками
		log.Println(l)
	}()

	//Настройка сервера
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("ошибка настройки сервера")
	}
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interseptors.UnaryPanicRecoveryInterceptor, //Обработка паник
		interseptors.XRequestID,                    //Создание ID запроса
		interseptors.LoggerInterseptor,             //Логирование запроса
	))
	spotAPI.RegisterSpotInstrumentServiceServer(grpcServer, spotService)

	fmt.Println("Сервер запущен на порту 8080...")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("ошибка работы сервера")
	}
	wg.Wait()
}
