package main

import (
	"Academy/gRPCServices/OrderService/pkg/interseptors"
	"Academy/gRPCServices/OrderService/pkg/methods"
	orderAPI "Academy/gRPCServices/Protobuf/Order"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8081") //Настройка сервера
	if err != nil {
		log.Fatal("Ошибка настройки сервиса", err)
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor( //регистрация интерсепторов
		interseptors.UnaryPanicRecoveryInterceptor,
		interseptors.XRequestID,
		interseptors.LoggerInterseptor,
	))
	orderServ := methods.NewOrderService()                     //Конструктор для OrderService
	orderAPI.RegisterOrderServiceServer(grpcServer, orderServ) //Регистрация методов

	fmt.Println("Сервер работает на порту 8081...")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Ошибка работы сервера:", err)
	}
}
