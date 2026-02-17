package postgres

import (
	postgresdto "Academy/gRPCServices/OrderService/internal/adapters/dto/postgres"
	"context"
	"database/sql"
	"errors"
	"time"
)

func (p *PostgresDB) AddMarketID(market_id int) (int, error) {
	dto := postgresdto.CreateMarketDTO(market_id)
	var id int
	err := p.QueryRow(context.Background(), `
	SELECT id FROM markets WHERE market_id = $1;
`, dto.Market_id).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		dto.Created_at = time.Now()
		//Добавить название маркета
		//Переделать контракт spotInstrument дляполучения названия рынка
		err = p.QueryRow(context.Background(), `
		INSERT INTO users(user_id,created_at) VALUES ($1,$2)
		`, dto.Market_id, dto.Created_at).Scan(&id)
		return id, nil
	}
	return 0, err
}
