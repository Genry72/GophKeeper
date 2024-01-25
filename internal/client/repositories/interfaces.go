package repositories

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

// IrepoSecrets хранилище секретов
type IrepoSecrets interface {
	// GetSecretTypes Получение типов секретов
	GetSecretTypes() []models.SecretType
	// CreateSecret Добавление секрета
	CreateSecret(secret models.SecretServerResponse) error
	// EditSecret Изменение имеющегося секрета
	EditSecret(secret any, secretID models.SecretID, typeID models.SecretTypeID) error
	// DeleteSecret Удаление секрета
	DeleteSecret(secretID models.SecretID, typeID models.SecretTypeID) error
	// GetSecretByID Получение секрета по его id и типу
	GetSecretByID(secretID models.SecretID, typeID models.SecretTypeID) (any, error)
	// GetSecretsByTypeID Получение всех секретов по переданному типу секрета
	GetSecretsByTypeID(typeID models.SecretTypeID) ([]any, error)
}

// IrepoSecretsSync Синхронизация информации с сервером
type IrepoSecretsSync interface {
	SetSecretTypes(src []models.SecretType)
	// SyncSecrets Прогрузка секретов с сервера
	SyncSecrets(ctx context.Context, src []models.SecretServerResponse) error
	// GetSecretTypes Получение тепов секретов, для прогрузки всех
	GetSecretTypes() []models.SecretType
}
