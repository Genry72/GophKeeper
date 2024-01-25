package secrets

import (
	"context"
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
