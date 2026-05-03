package apprunner

import (
	"context"
	"fmt"
	"log"
	"net/http"

	orderclient "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/order_client"
	spotclient "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/spot_client"
	userclient "github.com/DencCPU/gRPCServices/APIGetway/internal/adapters/user_client"
	"github.com/DencCPU/gRPCServices/APIGetway/internal/controller/gin"
	"github.com/DencCPU/gRPCServices/APIGetway/internal/usecase"
	"github.com/DencCPU/gRPCServices/Shared/breaker"
	"github.com/DencCPU/gRPCServices/Shared/logger"
	"github.com/DencCPU/gRPCServices/Shared/opentelemetry"
	"go.uber.org/zap"
)

func Apprunner() error {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	params := breaker.Params{
		Name:           "APIGetway",
		MaxRequest:     3,
		Interval:       10,
		Timeout:        5,
		MaxFailRequest: 5,
	}
	breaker := breaker.NewBreaker(logger, params)
	user_client, err := userclient.NewClient(breaker)
	if err != nil {
		logger.Error("user_client creation error",
			zap.Error(err),
		)
		return err
	}
	order_client, err := orderclient.NewClient(breaker)
	if err != nil {
		logger.Error("order_client creation error",
			zap.Error(err),
		)
		return err
	}

	spot_client, err := spotclient.NewClient(breaker)
	if err != nil {
		logger.Error("spot_client creation error",
			zap.Error(err),
		)
		return err
	}

	tracerProvider, err := opentelemetry.NewHttpTracer(context.Background(), "APIGetway", "localhost", "4318")
	if err != nil {
		return err
	}
	defer tracerProvider.Shutdown(context.Background())

	tracer := tracerProvider.Tracer("ApiGetway")

	service := usecase.NewService(user_client, order_client, spot_client, logger, tracer)
	api := gin.NewGinAPI(service)

	fmt.Println("Server is running")
	err = http.ListenAndServe(":8082", api.Router())
	if err != nil {
		logger.Error("server error:",
			zap.Error(err),
		)
		return err
	}
	return nil
}
