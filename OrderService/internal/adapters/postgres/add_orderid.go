package postgres

import (
	ordererrors "Academy/gRPCServices/OrderService/internal/domain/error"
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
	"database/sql"
	"errors"
	"math"
	"math/rand/v2"
)

func (p *PostgresDB) AddOrderID(newOrder order.Order, marketsID []int64) (int, error) {

	var foundMarket bool //Флаг, показывающий найден нужный рынок или нет

	for _, mId := range marketsID { //Проверка наличия нужного рынка
		if mId == newOrder.Market_id {
			foundMarket = true
			break
		}
	}

	if foundMarket != true {
		return 0, ordererrors.Avalible_markets
	}
	var orderID int
	for {
		orderID = rand.IntN(math.MaxInt64)
		var id int
		err := p.QueryRow(context.Background(), `
		SELECT order_id FROM orders_id WHERE order_id=$1;
		`, orderID).Scan(&id)
		if errors.Is(err, sql.ErrNoRows) {
			err = p.QueryRow(context.Background(), `
			INSERT INTO orders_id(order_id) VALUES ($1)
			`, orderID).Scan(&id)
			break
		}
	}
	return orderID, nil
}
