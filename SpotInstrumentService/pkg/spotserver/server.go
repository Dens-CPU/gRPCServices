package grpcserver

import (
	"Academy/gRPCServices/Shared/interseptors"
	spotconfig "Academy/gRPCServices/SpotInstrumentService/config"
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
	"net"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

// Структура сервера
type Server struct {
	*grpc.Server              //Сервер
	Listener     net.Listener //Настройки подключения
}

// Создание нового сервера
func New(redis *redisadapter.RedisDB, cfg spotconfig.Server) (*Server, error) {

	//Инициализация интерфеса listener
	host := cfg.Host
	port := cfg.Port
	lis, err := net.Listen(cfg.Network, host+port)
	if err != nil {
		return nil, err
	}

	//Регистрация интерсепторов на сервере
	newServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interseptors.UnaryPanicRecoveryInterceptor,
			interseptors.XRequestID,
			interseptors.LoggerInterseptor,
			redisadapter.RedisCacheInterceptor(redis, 10*time.Minute),
		),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	return &Server{newServer, lis}, nil
}
