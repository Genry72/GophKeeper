package secrets

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

func (s *SecretUc) GetSecretTypes(ctx context.Context) ([]models.SecretType, error) {
	return s.netClientSecrets.GetSecretTypes(ctx)
}
