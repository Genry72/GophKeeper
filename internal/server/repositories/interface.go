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

type ISecrets interface {
	GetSecretTypes(ctx context.Context) ([]models.SecretType, error)
	AddSecret(ctx context.Context, userID int64, secretTypeID int64,
		secretName string, secretContent []byte) (models.Secret, error)
	EditSecret(ctx context.Context,
		secretName string, secretID int64, secretContent []byte) (models.Secret, error)
	DeleteSecret(ctx context.Context, secretID int64) error
	GetSecretsBySecretTypeID(ctx context.Context,
		userID int64, typeID int64) ([]models.Secret, error)
	GetSecretByID(ctx context.Context,
		userID int64, secretID int64) (models.Secret, error)
}

type Repo struct {
	Users   IUsers
	Secrets ISecrets
	pgRepo  *postgres.PGStorage
}

func NewPostgresRepo(dsn string, log *zap.Logger) (*Repo, error) {
	pgRepo, err := postgres.NewPGStorage(dsn, log)
	if err != nil {
		return nil, fmt.Errorf("postgres.NewPGStorage: %w", err)
	}

	return &Repo{
		Users:   pgRepo.Users,
		Secrets: pgRepo.Secrets,
		pgRepo:  pgRepo,
	}, nil
}

func (r *Repo) Stop() {
	r.pgRepo.Stop()
}
