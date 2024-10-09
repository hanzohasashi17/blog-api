package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

type Database struct {
	Db *sql.DB
}

func New(storagePath string) (*Database, error) {
	const op = "database.sqlite.New"

	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := Migrate(db); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Database{Db: db}, nil
}

func Migrate(db *sql.DB) error {
	const op = "database.sqlite.Migrate"

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "sqlite", 
		driver,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("%s: %w", op, err)
	}
	
	return nil
}