package postgres

import (
	postgresdto "Academy/gRPCServices/OrderService/internal/adapters/dto/postgres"
	"context"
	"database/sql"
	"errors"
	"time"
)

func (p *PostgresDB) AddUserID(user_id int) (int, error) {
	dto := postgresdto.CreateUserDTO(user_id)
	var id int
	err := p.QueryRow(context.Background(), `
	SELECT id FROM users WHERE user_id = $1;
`, dto.User_id).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		dto.Created_at = time.Now()
		err = p.QueryRow(context.Background(), `
		INSERT INTO users(user_id,created_at) VALUES ($1,$2)
		`, dto.User_id, dto.Created_at).Scan(&id)
		return id, nil
	}
	return 0, err
}
