package postgres

import (
	postgresdto "Academy/gRPCServices/OrderService/internal/adapters/dto/postgres"
	"Academy/gRPCServices/OrderService/internal/domain/order"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func (p *PostgresDB) AddUserID(tx pgx.Tx, ctx context.Context, newOrder order.Order) (int, error) {
	//Инициализация DTO
	dto := postgresdto.CreateUserDTO(int(newOrder.User_id))
	dto.Created_at = time.Now()

	//Поиск пользователя с id
	var id int
	err := tx.QueryRow(ctx, `
		INSERT INTO users (user_id, created_at)
		VALUES ($1, $2)
		ON CONFLICT (user_id) DO UPDATE
		SET user_id = EXCLUDED.user_id
		RETURNING id
	`, dto.User_id, dto.Created_at).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
