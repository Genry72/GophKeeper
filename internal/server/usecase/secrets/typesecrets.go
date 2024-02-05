package secrets

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/server/models"
)

func (s *SecretsUsecase) GetSecretTypes(ctx context.Context) ([]models.SecretType, error) {
	return s.repo.Secrets.GetSecretTypes(ctx)
}
