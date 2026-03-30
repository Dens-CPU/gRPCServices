package postgres

import (
	"context"
	"time"

	postgresdto "github.com/DencCPU/gRPCServices/UserService/internal/adapters/dto/postgres"
	domainuser "github.com/DencCPU/gRPCServices/UserService/internal/domain/user"
	"github.com/google/uuid"
)

// Добавление нового пользоваткля
func (p *PostgresDB) AddUser(ctx context.Context, newUser domainuser.User) (string, string, error) {
	tx, err := p.Begin(ctx)
	if err != nil {
		return "", "", err
	}
	dto := postgresdto.NewUserDTO(newUser.Name, newUser.Email, newUser.Password, newUser.Role)
	dto.Created_at = time.Now()
	id := uuid.New()
	err = tx.QueryRow(ctx, `
	INSERT INTO users(id,name,email,password,role,created_at)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING id
	`, id, dto.Name, dto.Email, dto.Password, dto.Role, dto.Created_at).Scan(&dto.ID)

	if err != nil {
		tx.Rollback(ctx)
		return "", "", err
	}

	token, err := p.AddRefreshToken(tx, ctx, dto.ID)
	if err != nil {
		tx.Rollback(ctx)
		return "", "", err
	}
	tx.Commit(ctx)

	return dto.ID.String(), token, nil
}

// Замена пароля
func (p *PostgresDB) UpdatePassword(ctx context.Context, email, password string) error {
	dto := postgresdto.NewUpdatePassord(email, password)
	dto.Update_at = time.Now()

	_, err := p.Exec(ctx, `
	UPDATE users SET password = $1
	WHERE email = $2
	`, dto.Email, dto.Password)

	if err != nil {
		return err
	}
	return nil
}
