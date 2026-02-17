package postgres

import (
	"context"
	"errors"
)

func (p *PostgresDB) GetOrderID(orderID int) (int, error) {
	var id int
	err := p.QueryRow(context.Background(), `
	SELECT id FROM orders_id WHERE order_id = $1
	`, orderID).Scan(&id)
	if err != nil {
		return 0, errors.New("ID is not found")
	}
	return id, nil
}
