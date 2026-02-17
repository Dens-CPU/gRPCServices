package postgres

import (
	postgresdto "Academy/gRPCServices/OrderService/internal/adapters/dto/postgres"
	"Academy/gRPCServices/OrderService/internal/domain/order"

	"context"
	"time"
)

func (p *PostgresDB) AddOrderStorage(newOrder order.Order, orderID int) (string, error) {

	refOrderID, err := p.GetOrderID(orderID)
	if err != nil {
		return "", err
	}

	refUserID, err := p.AddUserID(int(newOrder.User_id))
	if err != nil {
		return "", err
	}

	refMarketID, err := p.AddMarketID(int(newOrder.Market_id))
	if err != nil {
		return "", err
	}

	dto := postgresdto.CreatOrderDTO(newOrder, refOrderID, refUserID, refMarketID)
	dto.Status = "created"
	dto.Created_at = time.Now()
	_, err = p.Exec(context.Background(), `
	INSERT INTO orders(user_id,market_id,order_type,price,quantity, status, order_id,created_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8);
	`,
		dto.Ref_User_Id,
		dto.Ref_Market_Id,
		dto.Order_type,
		dto.Price,
		dto.Quantity,
		dto.Status,
		dto.Ref_Order_Id,
		dto.Created_at)
	if err != nil {
		return "", err
	}

	return "create", nil
}
