package secrets

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Secrets struct {
	secretsClient proto.SecretClient
	log           *zap.Logger
}

func NewSecrets(grpcconn *grpc.ClientConn, log *zap.Logger) *Secrets {
	secretClient := proto.NewSecretClient(grpcconn)

	return &Secrets{
		secretsClient: secretClient,
		log:           log,
	}
}

// CreateSecret Сохранение секрета на сервере. Секрет в байтах
func (s *Secrets) CreateSecret(ctx context.Context, secretTypeID models.SecretTypeID,
	name models.SecretName, secretValue []byte) (models.SecretServerResponse, error) {
	// Отправка секрета на сервер
	resSecret, err := s.secretsClient.CreateSecret(ctx, &proto.CreateSecretRequest{
		Name:       string(name),
		SecretType: int64(secretTypeID),
		Data:       secretValue,
	})

	if err != nil {
		return models.SecretServerResponse{}, fmt.Errorf("s.secretsClient.CreateSecret: %w", err)
	}

	result := models.SecretServerResponse{
		ID:           models.SecretID(resSecret.Id),
		SecretTypeID: models.SecretTypeID(resSecret.SecretType),
		Name:         models.SecretName(resSecret.Name),
		Value:        resSecret.Content,
		CreatedAt:    resSecret.CreatedAt.AsTime(),
		UpdatedAt:    resSecret.UpdatedAt.AsTime(),
	}

	return result, nil
}

// EditSecret Редактирование секрета.
func (s *Secrets) EditSecret(ctx context.Context, secretID models.SecretID,
	name models.SecretName, secretValue []byte) (models.SecretServerResponse, error) {
	// Отправка секрета на сервер
	resSecret, err := s.secretsClient.EditSecret(ctx, &proto.EditSecretRequest{
		Id:   int64(secretID),
		Name: string(name),
		Data: secretValue,
	})

	if err != nil {
		return models.SecretServerResponse{}, fmt.Errorf("s.secretsClient.CreateSecret: %w", err)
	}

	result := models.SecretServerResponse{
		ID:           models.SecretID(resSecret.Id),
		SecretTypeID: models.SecretTypeID(resSecret.SecretType),
		Name:         models.SecretName(resSecret.Name),
		Value:        resSecret.Content,
		CreatedAt:    resSecret.CreatedAt.AsTime(),
		UpdatedAt:    resSecret.UpdatedAt.AsTime(),
	}

	return result, nil
}

// EditSecret Редактирование секрета
func (s *Secrets) DeleteSecret(ctx context.Context, secretID models.SecretID) error {
	// Отправка секрета на сервер
	_, err := s.secretsClient.DeleteSecret(ctx, &proto.DeleteSecretRequest{
		Id: int64(secretID),
	})
	if err != nil {
		return fmt.Errorf("s.secretsClient.DeleteSecret: %w", err)
	}

	return nil
}

func (s *Secrets) GetSecretsBySecretTypeID(ctx context.Context,
	secretTypeID models.SecretTypeID) ([]models.SecretServerResponse, error) {
	// Получение секретов с сервера
	resSecrets, err := s.secretsClient.GetSecretsByType(ctx, &proto.SecretsByTypeRequest{
		SecretType: int64(secretTypeID),
	})

	if err != nil {
		return nil, fmt.Errorf("s.secretsClient.GetSecretsByType: %w", err)
	}

	result := make([]models.SecretServerResponse, len(resSecrets.Secrets))

	for i := range resSecrets.Secrets {
		secret := models.SecretServerResponse{
			ID:           models.SecretID(resSecrets.Secrets[i].Id),
			SecretTypeID: models.SecretTypeID(resSecrets.Secrets[i].SecretType),
			Name:         models.SecretName(resSecrets.Secrets[i].Name),
			Value:        resSecrets.Secrets[i].Content,
			CreatedAt:    resSecrets.Secrets[i].CreatedAt.AsTime(),
			UpdatedAt:    resSecrets.Secrets[i].UpdatedAt.AsTime(),
		}
		result[i] = secret
	}

	return result, nil
}
