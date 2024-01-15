package postgres

import (
	"errors"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/repositories/postgres/users"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type PGStorage struct {
	Users *users.UsersRepo
}

func NewPGStorage(dsn string, log *zap.Logger) (*PGStorage, error) {
	if err := migration(dsn); err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Connect: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(1 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return &PGStorage{
		Users: users.NewUsersRepo(db, log),
	}, nil
}

func migration(dsn string) error {
	migration, err := migrate.New(
		"file://internal/server/repositories/postgres/migration",
		dsn)
	if err != nil {
		return fmt.Errorf("migrate.New: %w", err)
	}

	defer migration.Close()

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("m.Up: %w", err)
	}

	return nil
}
