package orderserver

import (
	ordrerconfig "Academy/gRPCServices/OrderService/config"
	"Academy/gRPCServices/Shared/interseptors"
	"fmt"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	Listener net.Listener
}

func New(cfg ordrerconfig.Server) (*Server, error) {

	host := cfg.Host
	port := cfg.Port
	dsn := fmt.Sprintf("%s:%d", host, port)
	lis, err := net.Listen(cfg.Network, dsn)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interseptors.UnaryPanicRecoveryInterceptor,
			interseptors.XRequestID,
			interseptors.LoggerInterseptor,
		), grpc.StatsHandler(otelgrpc.NewServerHandler()))

	return &Server{newServer, lis}, nil
}
