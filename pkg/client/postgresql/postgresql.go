package postgresql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Oleja123/dcaa-property/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, cfg config.DatabaseConfig) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	maxAttempts := cfg.MaxAttempts

	for ; maxAttempts > 0; maxAttempts -= 1 {
		ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.SecondsToConnect)*time.Second)

		conn, err := pgx.Connect(ctx, dsn)
		cancel()
		if err != nil {
			fmt.Printf("Попытка провалилась: %v\n", err)
			continue
		}
		return conn, nil
	}
	return nil, errors.New("не удалось подключиться к базе данных")
}
