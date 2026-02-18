package postgresdto

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"time"
)

type OrderDTO struct {
	Ref_User_Id   int
	Ref_Market_Id int
	Order_type    string
	Price         float64
	Quantity      int
	Status        string
	Ref_Order_Id  int
	Created_at    time.Time
}

type UsersDTO struct {
	User_id    int
	Created_at time.Time
}

type MarketDTO struct {
	Market_id  int
	Created_at time.Time
}

type Oreders_idDTO struct {
	Order_id   string
	Created_at time.Time
}

func CreatOrderDTO(order order.Order) *OrderDTO {
	input := OrderDTO{
		Order_type: order.Order_type,
		Price:      order.Price,
		Quantity:   int(order.Quantity),
	}
	return &input
}

func CreateUserDTO(user_id int) *UsersDTO {
	return &UsersDTO{User_id: user_id}
}

func CreateMarketDTO(market_id int) *MarketDTO {
	return &MarketDTO{Market_id: market_id}
}

func CreateOrders_idDTO(order_id string) *Oreders_idDTO {
	return &Oreders_idDTO{Order_id: order_id}
}
