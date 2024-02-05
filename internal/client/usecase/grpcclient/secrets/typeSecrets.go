package secrets

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/proto"
	"google.golang.org/grpc/status"
)

func (s *Secrets) GetSecretTypes(ctx context.Context) ([]models.SecretType, error) {
	req := &proto.SecretTypeRequest{}

	resultProto, err := s.secretsClient.GetSecretTypes(ctx, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			return nil, fmt.Errorf(e.Message())
		}
	}

	result := make([]models.SecretType, len(resultProto.SecretsType))

	for i := range resultProto.SecretsType {
		result[i] = models.SecretType{
			SecretTypeID:   models.SecretTypeID(resultProto.SecretsType[i].Id),
			SecretTypeName: models.SecretTypeName(resultProto.SecretsType[i].Name),
		}
	}

	return result, nil
}
