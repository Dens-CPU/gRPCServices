package orderserver

import (
	serverconfig "Academy/gRPCServices/OrderService/config/server"
	"Academy/gRPCServices/SpotInstrumentService/pkg/interseptors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	Listener net.Listener
}

func New() (*Server, error) {

	cfg, err := serverconfig.NewConfig()
	if err != nil {
		return nil, err
	}

	host := cfg.Server.Host
	port := cfg.Server.Port
	dsn := fmt.Sprintf("%s:%d", host, port)
	lis, err := net.Listen(cfg.Server.Network, dsn)
	if err != nil {
		return nil, err
	}

	interseptors := grpc.ChainUnaryInterceptor(
		interseptors.UnaryPanicRecoveryInterceptor,
		interseptors.XRequestID,
		interseptors.LoggerInterseptor,
	)
	newServer := grpc.NewServer(interseptors)
	return &Server{newServer, lis}, nil
}
