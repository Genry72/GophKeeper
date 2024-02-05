package secrets

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
)

func (s *SecretsRepo) GetSecretTypes(ctx context.Context) ([]models.SecretType, error) {
	query := "select id, name from secret_types"

	result := make([]models.SecretType, 0)

	if err := s.conn.SelectContext(ctx, &result, query); err != nil {
		return nil, fmt.Errorf("s.conn.GetContext: %w", err)
	}

	return result, nil
}
