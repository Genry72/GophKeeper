package usecase

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

// INetClientUsers Вызов методов сервера для работы с клиентами
type INetClientUsers interface {
	Register(ctx context.Context, username, password string) (string, error)
	Auth(ctx context.Context, username, password string) (string, error)
}

// InetClientSecrets Вызов методов сервера для работы с секретами.
type InetClientSecrets interface {
	GetSecretTypes(ctx context.Context) ([]models.SecretType, error)
	CreateSecret(ctx context.Context,
		secretID models.SecretTypeID, name models.SecretName, secretValue []byte) (models.SecretServerResponse, error)
	EditSecret(ctx context.Context, secretID models.SecretID,
		name models.SecretName, secretValue []byte) (models.SecretServerResponse, error)
	DeleteSecret(ctx context.Context, secretID models.SecretID) error
	GetSecretsBySecretTypeID(ctx context.Context,
		secretTypeID models.SecretTypeID) ([]models.SecretServerResponse, error)
}

// Iusers Регистрация и вход пользователя.
type Iusers interface {
	Auth(ctx context.Context, username, password string) error
	Register(ctx context.Context, username, password string) error
}

// ISecrets Управление секретами
type ISecrets interface {
	// GetSecretTypes Получение доступных типов секретов
	GetSecretTypes(ctx context.Context) ([]models.SecretType, error)
	// CreateSecret Добавление нового секрета
	CreateSecret(ctx context.Context,
		secretTypeID models.SecretTypeID, secretName models.SecretName, secretValue any) error
	// EditSecret Изменение секрета
	EditSecret(ctx context.Context,
		secretID models.SecretID, secretName models.SecretName, secretValue any) error
	// DeleteSecret Удаление секрета
	DeleteSecret(ctx context.Context, secretID models.SecretID, secretTypeID models.SecretTypeID) error
	// GetSecretBySecretTypeID Получение всех секретов по ID
	GetSecretBySecretTypeID(ctx context.Context, id models.SecretTypeID) ([]any, error)
}

// ISync синхронизация данных с сервером.
type ISync interface {
	StartSync(ctx context.Context) error
}
