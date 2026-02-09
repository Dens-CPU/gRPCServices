//В данном файле реализован контракт для сервиса OrderService

package ordermethods

import (
	ordermemory "Academy/gRPCServices/OrderService/pkg/memory"
	orderAPI "Academy/gRPCServices/Protobuf/Order"
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

// Итерфейс управляющий заказами
type OrdersRepository interface {
	AddOrder([]int64, ordermemory.Order) (int64, string, error) //Добавление заказа
	GetOrderState(key ordermemory.Key) (string, error)          //Получение статуса заказа
}

type OrderService struct {
	orderAPI.UnimplementedOrderServiceServer
	repo OrdersRepository
}

// Конструктор
func NewOrderService(repo OrdersRepository) *OrderService {
	return &OrderService{repo: repo}
}

// Создание заказа
func (o *OrderService) CreateOrder(ctx context.Context, req *orderAPI.CreateReq) (*orderAPI.CreateResp, error) {

	conn, err := grpc.NewClient(":8080", grpc.WithInsecure()) //Подклчение к SpotService
	if err != nil {
		return &orderAPI.CreateResp{}, fmt.Errorf("ошибка подключения к SpotService:%w", err)
	}
	defer conn.Close()

	sctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //Создание контекста для вызываемых методов SpotService
	defer cancel()

	spotClient := spotAPI.NewSpotInstrumentServiceClient(conn)            //Регитсрация нового клиента
	enableMarkets, err := spotClient.ViewMarket(sctx, &spotAPI.ViewReq{}) //Вызов метода SpotService
	if err != nil {
		return &orderAPI.CreateResp{}, fmt.Errorf("ошибка работы SpotService:%w", err)
	}
	marketsID := enableMarkets.EnableMarkets
	newOrder := ordermemory.Order{ //Запроса в структуру Order
		User_id:    req.UserId,
		Market_id:  req.MarketId,
		Order_type: req.OrderType,
		Price:      req.Price,
		Quantity:   req.Quantity,
	}

	id, stat, err := o.repo.AddOrder(marketsID, newOrder)
	if err != nil {
		return &orderAPI.CreateResp{}, err
	}

	return &orderAPI.CreateResp{OrderId: int64(id), Status: stat}, nil // Ответ сервера
}

// Получение статуса заказа
func (o *OrderService) GetOrderStatus(ctx context.Context, req *orderAPI.GetReq) (*orderAPI.GetResp, error) {
	orderKey := ordermemory.Key{Order_id: int(req.OrderId), User_id: req.UserId} //Формирование ключа
	status, err := o.repo.GetOrderState(orderKey)
	if err != nil {
		return &orderAPI.GetResp{}, err
	}
	return &orderAPI.GetResp{OrderStatus: status}, nil
}
