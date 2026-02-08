package memory

// Структура закакза
type Order struct {
	User_id    int64
	Market_id  int64
	Order_type string
	Price      float64
	Quantity   int64
	Status     string
}

type Key struct {
	User_id  int64
	Order_id int
}
