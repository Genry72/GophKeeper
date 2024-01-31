package secrets

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
)

func (s *SecretsRepo) AddSecret(ctx context.Context, userID int64, secretTypeID int64,
	secretName string, secretContent []byte) (models.Secret, error) {
	query := `INSERT INTO user_secrets (user_id, secret_type_id, name, data)
VALUES ($1, $2, $3, $4)
returning id, user_id, secret_type_id, name, data, created_at, updated_at, deleted_at
`

	var result models.Secret

	if err := s.conn.GetContext(ctx, &result, query, userID, secretTypeID, secretName, secretContent); err != nil {
		return result, fmt.Errorf("s.conn.GetContext: %w", err)
	}

	return result, nil
}

// EditSecret Редактирование секрета
func (s *SecretsRepo) EditSecret(ctx context.Context,
	secretName string, secretID int64, secretContent []byte) (models.Secret, error) {
	query := `UPDATE user_secrets
SET 
    name = $1,
    data = $2,
    updated_at = now()
WHERE id = $3
returning id, user_id, secret_type_id, name, data, created_at, updated_at, deleted_at
`

	var result models.Secret

	if err := s.conn.GetContext(ctx, &result, query, secretName, secretContent, secretID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, models.ErrUserNotFound
		}
		return result, fmt.Errorf("s.conn.GetContext: %w", err)
	}

	return result, nil
}

// DeleteSecret Удаление секрета.
func (s *SecretsRepo) DeleteSecret(ctx context.Context, secretID int64) error {
	query := `UPDATE user_secrets
SET 
    updated_at = now(),
	deleted_at = now()
WHERE id = $1
`

	if _, err := s.conn.ExecContext(ctx, query, secretID); err != nil {
		return fmt.Errorf("s.conn.ExecContext: %w", err)
	}

	return nil
}

func (s *SecretsRepo) GetSecretsBySecretTypeID(ctx context.Context,
	userID int64, typeID int64) ([]models.Secret, error) {
	query := `
select id,
       user_id,
       secret_type_id,
       name,
       data,
       created_at,
       updated_at,
       deleted_at
from user_secrets
where deleted_at is null
and user_id = $1
and secret_type_id = $2
`

	var result []models.Secret

	if err := s.conn.SelectContext(ctx, &result, query, userID, typeID); err != nil {
		return nil, fmt.Errorf("s.conn.SelectContext: %w", err)
	}

	return result, nil
}

// GetSecretByID Получение секрета по ID
func (s *SecretsRepo) GetSecretByID(ctx context.Context,
	userID int64, secretID int64) (models.Secret, error) {
	query := `
select id,
       user_id,
       secret_type_id,
       name,
       data,
       created_at,
       updated_at,
       deleted_at
from user_secrets
where 
    id = $1
and deleted_at is null
and user_id = $2
`

	var result models.Secret

	if err := s.conn.GetContext(ctx, &result, query, secretID, userID); err != nil {
		return result, fmt.Errorf("s.conn.GetContext: %w", err)
	}

	return result, nil
}
