//В данном файле реализован контракт для сервиса OrderService

package methods

import (
	controlorderstat "Academy/gRPCServices/OrderService/pkg/controlOrderStat"
	"Academy/gRPCServices/OrderService/pkg/memory"
	orderAPI "Academy/gRPCServices/Protobuf/Order"
	spotAPI "Academy/gRPCServices/Protobuf/Spot"
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

type OrderService struct {
	orderAPI.UnimplementedOrderServiceServer
	Storage map[memory.Key]*memory.Order
	orders  map[int]struct{}
}

// Конструктор
func NewOrderService() *OrderService {
	return &OrderService{Storage: make(map[memory.Key]*memory.Order), orders: make(map[int]struct{})}
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

	newOrder := memory.Order{ //Запроса в структуру Order
		User_id:    req.UserId,
		Market_id:  req.MarketId,
		Order_type: req.OrderType,
		Price:      req.Price,
		Quantity:   req.Quantity,
	}

	var foundMarket bool //Флаг, показывающий найден нужный рынок или нет

	for _, mId := range enableMarkets.EnableMarkets { //Проверка наличия нужного рынка
		if mId == newOrder.Market_id {
			foundMarket = true
		}
	}
	if foundMarket != true {
		return &orderAPI.CreateResp{}, errors.New("Рынок недоступен")
	}

	var orderId int //ID нового заказа
	for {
		id := rand.Intn(math.MaxInt64) //Создание ID заказа
		if _, exist := o.orders[id]; !exist {
			o.orders[id] = struct{}{}
			orderId = id
			break
		}
	}

	key := memory.Key{User_id: newOrder.User_id, Order_id: orderId} //Создание ключа для in-memory хранилища
	o.Storage[key] = &newOrder
	o.Storage[key].Status = "created" //Сохранение заказа в памяти

	controlorderstat.ControlStat(key, o.Storage)
	return &orderAPI.CreateResp{OrderId: int64(orderId), Status: "created"}, nil // Ответ сервера
}

func (o *OrderService) GetOrderStatus(ctx context.Context, req *orderAPI.GetReq) (*orderAPI.GetResp, error) {
	orderKey := memory.Key{Order_id: int(req.OrderId), User_id: req.UserId}
	if _, exist := o.Storage[orderKey]; !exist {
		return &orderAPI.GetResp{}, errors.New("Заказа не существует")
	}
	status := o.Storage[orderKey].Status
	return &orderAPI.GetResp{OrderStatus: status}, nil
}
