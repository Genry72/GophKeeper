package secrets

import (
	"github.com/Genry72/GophKeeper/internal/server/repositories"
	"go.uber.org/zap"
)

type SecretsUsecase struct {
	repo *repositories.Repo
	log  *zap.Logger
}

func NewSecretsUsecase(repo *repositories.Repo, log *zap.Logger) *SecretsUsecase {
	return &SecretsUsecase{
		repo: repo,
		log:  log,
	}
}
