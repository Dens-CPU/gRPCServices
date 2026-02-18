package postgres

import (
	postgresdto "Academy/gRPCServices/OrderService/internal/adapters/dto/postgres"
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func (p *PostgresDB) AddMarketID(tx pgx.Tx, ctx context.Context, newOrder order.Order) (int, error) {
	//Инициализация DTO
	dto := postgresdto.CreateMarketDTO(int(newOrder.Market_id))
	dto.Created_at = time.Now()

	//Добавление маркета
	var id int
	err := tx.QueryRow(ctx, `
		INSERT INTO markets (market_id, created_at)
		VALUES ($1, $2)
		ON CONFLICT (market_id) DO UPDATE
		SET market_id = EXCLUDED.market_id
		RETURNING id
	`, dto.Market_id, dto.Created_at).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
