package order

import "github.com/shopspring/decimal"

// Структура закакза
type Order struct {
	User_id    int64
	Market_id  int64
	Order_type string
	Price      decimal.Decimal
	Quantity   int64
	Status     string
}

type Key struct {
	User_id  int64
	Order_id string
}

type Market struct {
	ID   int64
	Name string
}
