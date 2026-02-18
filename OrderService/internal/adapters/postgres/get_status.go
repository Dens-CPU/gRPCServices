package postgres

import (
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
)

func (p *PostgresDB) GetOrderState(ctx context.Context, key order.Key) (string, error) {
	var status string

	err := p.QueryRow(ctx, `
	SELECT status 
	FROM orders
	JOIN users ON orders.user_id = users.id
	JOIN orders_id ON orders.order_id = orders_id.id
	WHERE users.user_id = $1
	AND orders_id.order_id = $2
`, key.User_id, key.Order_id).Scan(&status)
	if err != nil {
		return "", err
	}
	return status, nil
}
