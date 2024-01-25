package secrets

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
)

func (s *SecretsUsecase) AddSecret(ctx context.Context, secretTypeID models.SecretTypeID,
	secretName string, secretContent []byte) (models.Secret, error) {
	userID, ok := ctx.Value(models.CtxUserID{}).(int64)
	if !ok {
		return models.Secret{}, fmt.Errorf("not UserID in context")
	}

	return s.repo.Secrets.AddSecret(ctx, userID, int64(secretTypeID), secretName, secretContent)
}

func (s *SecretsUsecase) GetSecretsBySecretTypeID(ctx context.Context,
	secretTypeID models.SecretTypeID) ([]models.Secret, error) {
	userID, ok := ctx.Value(models.CtxUserID{}).(int64)
	if !ok {
		return nil, fmt.Errorf("not UserID in context")
	}

	return s.repo.Secrets.GetSecretsBySecretTypeID(ctx, userID, int64(secretTypeID))
}
