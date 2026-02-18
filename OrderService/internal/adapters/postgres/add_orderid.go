package postgres

import (
	postgresdto "Academy/gRPCServices/OrderService/internal/adapters/dto/postgres"
	ordererrors "Academy/gRPCServices/OrderService/internal/domain/error"
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (p *PostgresDB) AddOrderID(tx pgx.Tx, ctx context.Context, newOrder order.Order, marketsID []int64) (int, string, error) {

	var foundMarket bool //Флаг, показывающий найден нужный рынок или нет

	//Проверка наличия нужного рынка
	for _, mId := range marketsID {
		if mId == newOrder.Market_id {
			foundMarket = true
			break
		}
	}
	if foundMarket != true {
		return 0, "", ordererrors.Avalible_markets
	}

	//Присвоение нового id заказа
	var orderID string
	orderID = uuid.New().String()

	//Ининциализация DTO
	dto := postgresdto.CreateOrders_idDTO(orderID)
	dto.Created_at = time.Now()

	//Запись в БД
	var id int
	err := tx.QueryRow(context.Background(), ` 
	INSERT INTO orders_id(order_id,created_at) 
	VALUES ($1,$2)
	RETURNING id
	`, dto.Order_id, dto.Created_at).Scan(&id)
	if err != nil {
		return 0, "", err
	}

	return id, orderID, nil
}
