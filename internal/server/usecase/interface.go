package usecase

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/server/models"
	"github.com/Genry72/GophKeeper/internal/server/repositories"
	"github.com/Genry72/GophKeeper/internal/server/usecase/secrets"
	"github.com/Genry72/GophKeeper/internal/server/usecase/users"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"go.uber.org/zap"
)

// Iusers Работа с пользователями
type Iusers interface {
	RegisterUser(ctx context.Context, login, pass string) (int64, error)
	AuthUser(ctx context.Context, login, pass string) (string, error)
}

// ISecrets Работа с секретами
type ISecrets interface {
	GetSecretTypes(ctx context.Context) ([]models.SecretType, error)
	AddSecret(ctx context.Context, secretTypeID models.SecretTypeID, secretName string,
		secretContent []byte) (models.Secret, error)
	GetSecretsBySecretTypeID(ctx context.Context,
		secretTypeID models.SecretTypeID) ([]models.Secret, error)
}

type Usecase struct {
	Users   Iusers
	Secrets ISecrets
	log     *zap.Logger
}

func NewUsecase(repo *repositories.Repo, jwtService *jwttoken.Service, log *zap.Logger) *Usecase {
	return &Usecase{
		Users:   users.NewUsersUsecase(repo, jwtService, log),
		Secrets: secrets.NewSecretsUsecase(repo, log),
		log:     log,
	}
}
