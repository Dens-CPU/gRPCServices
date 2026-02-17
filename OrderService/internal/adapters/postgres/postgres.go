package postgres

import (
	postgresconfig "Academy/gRPCServices/OrderService/config/postgres"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	*pgxpool.Pool
}

func NewDB(ctx context.Context) (*PostgresDB, error) {
	cfg, err := postgresconfig.NewConfig()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Name,
		cfg.Postgres.Sslmode,
	)
	fmt.Println(dsn)
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres database is unavailable:%w", err) //Нужно ли эту ошибку отнести в доменную область?
	}

	// Проверка соединения
	conn, err := db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot acquire connection from pool:%w", err)
	}
	defer conn.Release()

	return &PostgresDB{db}, nil
}
