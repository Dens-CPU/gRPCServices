package userserver

import (
	"net"

	"github.com/DencCPU/gRPCServices/Shared/interseptors"
	userconfig "github.com/DencCPU/gRPCServices/UserService/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	Listener net.Listener
}

func NewServer(cfg userconfig.Server, logger *zap.Logger) (*Server, error) {
	//Инициализация интерфеса listener
	host := cfg.Host
	port := cfg.Port
	lis, err := net.Listen(cfg.Network, host+":"+port)
	if err != nil {
		return nil, err
	}

	//Регистрация интерсепторов на сервере
	newServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interseptors.UnaryPanicRecoveryInterceptor(logger),
			interseptors.XRequestID,
			interseptors.LoggerInterseptor(logger),
		),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	return &Server{newServer, lis}, nil
}
