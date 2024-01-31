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

func (s *SecretsUsecase) EditSecret(ctx context.Context,
	secretName string, secretID int64, secretContent []byte) (models.Secret, error) {
	userID, ok := ctx.Value(models.CtxUserID{}).(int64)
	if !ok {
		return models.Secret{}, fmt.Errorf("not UserID in context")
	}

	if _, err := s.repo.Secrets.GetSecretByID(ctx, userID, secretID); err != nil {
		return models.Secret{}, err
	}

	return s.repo.Secrets.EditSecret(ctx, secretName, secretID, secretContent)
}

func (s *SecretsUsecase) DeleteSecret(ctx context.Context, secretID int64) error {
	userID, ok := ctx.Value(models.CtxUserID{}).(int64)
	if !ok {
		return fmt.Errorf("not UserID in context")
	}

	if _, err := s.repo.Secrets.GetSecretByID(ctx, userID, secretID); err != nil {
		return err
	}

	return s.repo.Secrets.DeleteSecret(ctx, secretID)
}

func (s *SecretsUsecase) GetSecretsBySecretTypeID(ctx context.Context,
	secretTypeID models.SecretTypeID) ([]models.Secret, error) {
	userID, ok := ctx.Value(models.CtxUserID{}).(int64)
	if !ok {
		return nil, fmt.Errorf("not UserID in context")
	}

	return s.repo.Secrets.GetSecretsBySecretTypeID(ctx, userID, int64(secretTypeID))
}

func (s *SecretsUsecase) GetSecretByID(ctx context.Context,
	secretID int64) (models.Secret, error) {
	userID, ok := ctx.Value(models.CtxUserID{}).(int64)
	if !ok {
		return models.Secret{}, fmt.Errorf("not UserID in context")
	}

	return s.repo.Secrets.GetSecretByID(ctx, userID, secretID)
}
