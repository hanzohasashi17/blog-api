package postgres

import (
	"context"
	"fmt"

	migrator "github.com/cybertec-postgresql/pgx-migrator"
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

	if err := migrate(pool); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return pool, nil
}

func migrate(pool *pgxpool.Pool) error {
	op := "database.postgres.migrate"

	migrateQuery := "CREATE TABLE posts (id SERIAL PRIMARY KEY, title TEXT NOT NULL, content TEXT NOT NULL, author TEXT NOT NULL, created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW())"

	m, err := migrator.New(
        migrator.Migrations(
            &migrator.MigrationNoTx{
                Name: "Create table posts",
                Func: func(ctx context.Context, migrator migrator.PgxIface) error {
                    _, err := migrator.Exec(ctx, migrateQuery)
                    return err
                },
            },
        ),
    )
    if err != nil {
        return fmt.Errorf("%s: %w", op, err)
    }

	if err := m.Migrate(context.Background(), pool); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}