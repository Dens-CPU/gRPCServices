package grpcserver

import (
	serverconfig "Academy/gRPCServices/SpotInstrumentService/config/server"
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
	"Academy/gRPCServices/SpotInstrumentService/pkg/interseptors"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	Listener net.Listener
}

func New(redis *redisadapter.RedisDB) (*Server, error) {

	cfg, err := serverconfig.NewConfig()
	if err != nil {
		return nil, err
	}

	host := cfg.Server.Host
	port := cfg.Server.Port
	lis, err := net.Listen(cfg.Server.Network, host+port)
	if err != nil {
		return nil, err
	}

	interseptors := grpc.ChainUnaryInterceptor(
		interseptors.UnaryPanicRecoveryInterceptor,
		interseptors.XRequestID,
		interseptors.LoggerInterseptor,
		interseptors.RedisCacheInterceptor(redis, 10*time.Minute),
	)
	newServer := grpc.NewServer(interseptors)
	return &Server{newServer, lis}, nil
}
