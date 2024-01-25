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
	sync             usecase.ISync             // Синхронизация данных с сервером
	log              *zap.Logger
}

func NewSecretUc(netClient usecase.InetClientSecrets,
	localRepo repositories.IrepoSecrets, sync usecase.ISync, log *zap.Logger) *SecretUc {

	return &SecretUc{
		netClientSecrets: netClient,
		localRepo:        localRepo,
		sync:             sync,
		log:              log,
	}
}

// CreateSecret Отправляет секрет на сервер, полученный секрет от сервера добавляет в локальное хранилище
func (s *SecretUc) CreateSecret(ctx context.Context,
	secretTypeID models.SecretTypeID, secretName models.SecretName, secretValue any) error {
	// Отправка секрета на сервер
	secretByte, err := json.Marshal(secretValue)
	if err != nil {
		return fmt.Errorf("json.Marshal secretValue: %w", err)
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

func (s *SecretUc) GetSecretBySecretTypeID(ctx context.Context, id models.SecretTypeID) ([]any, error) {
	return s.localRepo.GetSecretsByTypeID(id)
}
