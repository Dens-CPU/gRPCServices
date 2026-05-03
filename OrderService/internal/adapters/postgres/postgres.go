package postgres

import (
	"context"
	"fmt"
	"sync"

	orderconfig "github.com/DencCPU/gRPCServices/OrderService/config"
	orderdomain "github.com/DencCPU/gRPCServices/OrderService/internal/domain/order"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Notify interface {
	AddNewState(userId string, orderId string, statCh chan string)
}

type PostgresDB struct {
	*pgxpool.Pool
	Notify
	controlOrderChan chan orderdomain.OrderInfo

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewDB(ctx context.Context, cfg orderconfig.Postgres, notify Notify) (*PostgresDB, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Sslmode,
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

	dbCtx, dbCancel := context.WithCancel(ctx)

	return &PostgresDB{db, notify, make(chan orderdomain.OrderInfo, cfg.ControlChanSize), dbCtx, dbCancel, sync.WaitGroup{}}, nil
}
