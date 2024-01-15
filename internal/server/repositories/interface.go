package repositories

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
	"github.com/Genry72/GophKeeper/internal/server/repositories/postgres"
	"go.uber.org/zap"
)

// IUsers Работа с пользователями
type IUsers interface {
	// Register Добавление нового пользователя в базу
	Register(ctx context.Context, login, encryptPass string) (int64, error)
	// FindByLogin поиск по логину
	FindByLogin(ctx context.Context, login string) (*models.Users, bool, error)
	// FindByID поиск по id пользователя
	FindByID(ctx context.Context, id int64) (*models.Users, bool, error)
}

type Repo struct {
	Users IUsers
}

func NewPostgresRepo(dsn string, log *zap.Logger) (*Repo, error) {
	pgRepo, err := postgres.NewPGStorage(dsn, log)
	if err != nil {
		return nil, fmt.Errorf("postgres.NewPGStorage: %w", err)
	}
	return &Repo{
		Users: pgRepo.Users,
	}, nil
}
