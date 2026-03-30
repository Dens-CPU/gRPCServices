package postgres

import (
	"context"
	"time"

	jwtdto "github.com/DencCPU/gRPCServices/UserService/internal/adapters/dto/tokens"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// Добавление refresh token в БД
func (p *PostgresDB) AddRefreshToken(tx pgx.Tx, ctx context.Context, user_id uuid.UUID) (string, error) {

	//Создание нового токена
	refreshToken := jwtdto.NewRefreshToken(user_id)

	var token = refreshToken.Token
	refreshToken.Created_at = time.Now()

	_, err := tx.Exec(ctx, `
	INSERT INTO refresh_tokens(token,expire_at,is_revoked,user_id,created_at)
	VALUES($1,$2,$3,$4,$5) 
	RETURNING token 
	`,
		refreshToken.Token,
		refreshToken.Expire_at,
		refreshToken.IsRevoked,
		refreshToken.UserId,
		refreshToken.Created_at,
	)
	if err != nil {
		return "", err
	}
	return token, nil
}
