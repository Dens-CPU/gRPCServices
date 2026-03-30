package postgresdto

import (
	"time"

	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	"github.com/shopspring/decimal"
)

type OrderDTO struct {
	Ref_User_Id   int
	Ref_Market_Id int
	Order_type    string
	Price         decimal.Decimal
	Quantity      int
	Status        string
	Ref_Order_Id  int
	Created_at    time.Time
}

type UsersDTO struct {
	User_id    string
	Created_at time.Time
}

type MarketDTO struct {
	Market_id   string
	Market_name string
	Created_at  time.Time
}

type Oreders_idDTO struct {
	Order_id   string
	Created_at time.Time
}

func CreatOrderDTO(order orderdomain.Order) *OrderDTO {
	input := OrderDTO{
		Order_type: order.Order_type,
		Price:      order.Price,
		Quantity:   int(order.Quantity),
	}
	return &input
}

func CreateUserDTO(user_id string) *UsersDTO {
	return &UsersDTO{User_id: user_id}
}

func CreateMarketDTO(market_id, market_name string) *MarketDTO {
	return &MarketDTO{Market_id: market_id, Market_name: market_name}
}

func CreateOrders_idDTO(order_id string) *Oreders_idDTO {
	return &Oreders_idDTO{Order_id: order_id}
}
