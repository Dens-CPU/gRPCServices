package main

import (
	"fmt"
	"log"
	"net/http"

	userservice "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/user_service"
	"github.com/DencCPU/gRPCServices/APIGetway/internal/controller/gin"
	"github.com/DencCPU/gRPCServices/APIGetway/internal/usecase"
	"github.com/DencCPU/gRPCServices/Shared/logger"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	client, err := userservice.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	service := usecase.NewService(client, logger)
	api := gin.NewGinAPI(service)
	fmt.Println("Сервер запущен")
	err = http.ListenAndServe(":8080", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
