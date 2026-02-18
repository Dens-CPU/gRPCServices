package memory

import (
	ordererrors "Academy/gRPCServices/OrderService/internal/domain/error"
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"math"
	"math/rand"
	"strconv"
)

func (s *Storage) AddOrderID(newOrder order.Order, marketsID []int64) (string, error) {

	var foundMarket bool //Флаг, показывающий найден нужный рынок или нет

	for _, mId := range marketsID { //Проверка наличия нужного рынка
		if mId == newOrder.Market_id {
			foundMarket = true
			break
		}
	}

	if foundMarket != true {
		return "", ordererrors.Avalible_markets
	}
	var orderId int //ID нового заказа
	for {
		id := rand.Intn(math.MaxInt64) //Создание ID заказа
		if _, exist := s.ordersID[id]; !exist {
			s.ordersID[id] = struct{}{}
			orderId = id
			break
		}
	}

	return strconv.Itoa(orderId), nil
}
