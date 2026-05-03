package postgres

import (
	"context"
	"time"

	sharederrors "github.com/DencCPU/gRPCServices/Shared/errors"
	tokensdto "github.com/DencCPU/gRPCServices/UserService/internal/adapters/dto/tokens"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// Adding a new refresh token to database
func (p *PostgresDB) AddRefreshToken(tx pgx.Tx, ctx context.Context, userId uuid.UUID) (string, error) {

	//Create a new refresh token
	refreshToken := tokensdto.NewRefreshToken(userId)

	var token = refreshToken.Token
	refreshToken.CreatedAt = time.Now()

	_, err := tx.Exec(ctx, `
	INSERT INTO refresh_tokens(token,expire_at,is_revoked,user_id,created_at)
	VALUES($1,$2,$3,$4,$5) 
	RETURNING token 
	`,
		refreshToken.Token,
		refreshToken.ExpireAt,
		refreshToken.IsRevoked,
		refreshToken.UserId,
		refreshToken.CreatedAt,
	)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Update token data
func (p *PostgresDB) UpdateRefreshToken(ctx context.Context, token string) (string, error) {

	tx, err := p.Begin(ctx)
	if err != nil {
		return "", err
	}

	var rToken tokensdto.RefreshToken
	err = tx.QueryRow(ctx, `
	  SELECT id, token, expire_at, is_revoked 
        FROM refresh_tokens
        WHERE token = $1
	`, token).Scan(
		&rToken.ID,
		&rToken.Token,
		&rToken.ExpireAt,
		&rToken.IsRevoked)
	if err != nil {
		return "", err
	}

	if rToken.IsRevoked {
		return "", sharederrors.UserBlocked
	}

	if time.Now().After(rToken.ExpireAt) {
		return "", sharederrors.ReAutentification
	}

	rToken.Token = uuid.NewString()
	rToken.UpdateAt = time.Now()

	_, err = tx.Exec(ctx, `
	UPDATE refresh_tokens SET 
	token=$1,
	update_at =$2
	WHERE id = $3
	`,
		rToken.Token,
		rToken.UpdateAt,
		rToken.ID)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)
	return rToken.Token, tx.Commit(ctx)
}

func (p *PostgresDB) UpdateExpireAt(ctx context.Context, userId string) (string, error) {

	tx, err := p.Begin(ctx)
	if err != nil {
		return "", err
	}

	var isRevoked bool
	err = tx.QueryRow(ctx, `
	SELECT is_revoked 
        FROM refresh_tokens
        WHERE user_id = $1
	`, userId).Scan(&isRevoked)

	if isRevoked {
		return "", sharederrors.UserBlocked
	}

	expire_at := time.Now().Add(5 * 24 * time.Hour)
	token := uuid.NewString()
	update_at := time.Now()

	_, err = tx.Exec(ctx, `
	UPDATE refresh_tokens SET
	token=$1,
	expire_at=$2,
	update_at=$3
	WHERE user_id = $4
	`,
		token,
		expire_at,
		update_at,
		userId,
	)

	if err != nil {
		return "", err
	}

	defer tx.Rollback(ctx)
	return token, tx.Commit(ctx)
}
