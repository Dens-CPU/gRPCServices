package grpcserver

import (
	serverconfig "Academy/gRPCServices/SpotInstrumentService/config/server"
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
	"Academy/gRPCServices/SpotInstrumentService/pkg/interseptors"
	"net"
	"time"

	"google.golang.org/grpc"
)

// Структура сервера
type Server struct {
	*grpc.Server              //Сервер
	Listener     net.Listener //Настройки подключения
}

// Создание нового сервера
func New(redis *redisadapter.RedisDB) (*Server, error) {

	//Получение параметров подключения из конфига
	cfg, err := serverconfig.NewConfig()
	if err != nil {
		return nil, err
	}

	//Инициализация интерфеса listener
	host := cfg.Server.Host
	port := cfg.Server.Port
	lis, err := net.Listen(cfg.Server.Network, host+port)
	if err != nil {
		return nil, err
	}

	//Регистрация интерсепторов на сервере
	interseptors := grpc.ChainUnaryInterceptor(
		interseptors.UnaryPanicRecoveryInterceptor,                //Обработка паник
		interseptors.XRequestID,                                   //Формирование ID запроса
		interseptors.LoggerInterseptor,                            //Логирование заросов
		interseptors.RedisCacheInterceptor(redis, 10*time.Minute), //Сохранение запросов в кэш
	)
	newServer := grpc.NewServer(interseptors)

	return &Server{newServer, lis}, nil
}
