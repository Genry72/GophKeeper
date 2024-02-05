package secrets

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/internal/client/repositories"
	"github.com/Genry72/GophKeeper/internal/client/usecase"
	"go.uber.org/zap"
)

type SecretUc struct {
	netClientSecrets usecase.InetClientSecrets // Сетевой клиент для обмена с севрером
	localRepo        repositories.IrepoSecrets // Локальное хранение секретов
	log              *zap.Logger
}

func NewSecretUc(netClient usecase.InetClientSecrets,
	localRepo repositories.IrepoSecrets, log *zap.Logger) *SecretUc {

	return &SecretUc{
		netClientSecrets: netClient,
		localRepo:        localRepo,
		log:              log,
	}
}

// CreateSecret Отправляет секрет на сервер, полученный секрет от сервера добавляет в локальное хранилище
func (s *SecretUc) CreateSecret(ctx context.Context,
	secretTypeID models.SecretTypeID, secretName models.SecretName, secretValue any) error {

	var (
		secretByte []byte
		err        error
	)

	// Отправка секрета на сервер
	switch sv := secretValue.(type) {
	case models.SecretBinaryValue:
		secretByte = sv
	default:
		secretByte, err = json.Marshal(secretValue)
		if err != nil {
			return fmt.Errorf("json.Marshal secretValue: %w", err)
		}
	}

	resSecret, err := s.netClientSecrets.CreateSecret(ctx, secretTypeID, secretName, secretByte)
	if err != nil {
		return fmt.Errorf("s.netClientSecrets.CreateSecret: %w", err)
	}

	// Сохранение секрета в локальном хранилище
	if err := s.localRepo.CreateSecret(resSecret); err != nil {
		return fmt.Errorf("s.localRepo.CreateSecret: %w", err)
	}

	return nil
}

// EditSecret Изменение секрета
func (s *SecretUc) EditSecret(ctx context.Context,
	secretID models.SecretID, secretName models.SecretName, secretValue any) error {

	// Отправка секрета на сервер
	secretByte, err := json.Marshal(secretValue)
	if err != nil {
		return fmt.Errorf("json.Marshal secretValue: %w", err)
	}

	resSecret, err := s.netClientSecrets.EditSecret(ctx, secretID, secretName, secretByte)
	if err != nil {
		return fmt.Errorf("s.netClientSecrets.CreateSecret: %w", err)
	}

	// Изменение секрета в локальном хранилище
	if err := s.localRepo.EditSecret(resSecret, secretID, resSecret.SecretTypeID); err != nil {
		return fmt.Errorf("s.localRepo.CreateSecret: %w", err)
	}

	return nil
}

// DeleteSecret Удаление секрета.
func (s *SecretUc) DeleteSecret(ctx context.Context, secretID models.SecretID, secretTypeID models.SecretTypeID) error {
	// удаление на сервере
	if err := s.netClientSecrets.DeleteSecret(ctx, secretID); err != nil {
		return fmt.Errorf("s.netClientSecrets.DeleteSecret: %w", err)
	}
	// Удаление в локальном хранилище
	return s.localRepo.DeleteSecret(secretID, secretTypeID)
}

func (s *SecretUc) GetSecretBySecretTypeID(ctx context.Context, id models.SecretTypeID) ([]any, error) {
	return s.localRepo.GetSecretsByTypeID(id)
}
