package postgresql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Oleja123/dcaa-property/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(ctx context.Context, cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	maxAttempts := cfg.MaxAttempts

	for ; maxAttempts > 0; maxAttempts -= 1 {
		ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.SecondsToConnect)*time.Second)

		pool, err := pgxpool.New(ctx, dsn)
		cancel()
		if err != nil {
			fmt.Printf("Попытка провалилась: %v\n", err)
			continue
		}
		return pool, nil
	}
	return nil, errors.New("не удалось подключиться к базе данных")
}
