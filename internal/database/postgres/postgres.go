package postgres

import (
	"context"
	"fmt"

	"github.com/hanzohasashi17/blog-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg config.Database) (*pgxpool.Pool, error) {
	op := "database.postgres.New"

	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%v/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return pool, nil
}