package orderdomain

import "github.com/shopspring/decimal"

// Структура закакза
type Order struct {
	User_id    string
	Market_id  string
	Order_type string
	Price      decimal.Decimal
	Quantity   int64
	Status     string
}

type Key struct {
	User_id  string //uuid
	Order_id string
}

type Market struct {
	ID   string
	Name string
}
