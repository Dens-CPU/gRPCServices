package orderserver

import (
	"fmt"
	"net"

	orderconfig "github.com/DencCPU/gRPCServices/OrderService/config"
	"github.com/DencCPU/gRPCServices/Shared/interseptors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	Listener net.Listener
}

func New(cfg orderconfig.Server, logger *zap.Logger) (*Server, error) {

	host := cfg.Host
	port := cfg.Port
	dsn := fmt.Sprintf("%s:%d", host, port)
	lis, err := net.Listen(cfg.Network, dsn)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interseptors.UnaryPanicRecoveryInterceptor(logger),
			interseptors.XRequestID,
			interseptors.LoggerInterseptor(logger),
		), grpc.StatsHandler(otelgrpc.NewServerHandler()))

	return &Server{newServer, lis}, nil
}
