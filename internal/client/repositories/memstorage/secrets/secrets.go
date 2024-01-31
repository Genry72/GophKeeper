package secrets

import (
	"context"
	"errors"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/pkg/helper"
	"sort"
	"sync"
)

type Secrets struct {
	typeSecrets     []models.SecretType
	secretLogPass   map[models.SecretID]models.SecretLogPass
	secretsText     map[models.SecretID]models.SecretText
	secretsBinary   map[models.SecretID]models.SecretBinary
	secretsBankCard map[models.SecretID]models.SecretBankCard
	mx              sync.RWMutex
}

func NewSecrets() *Secrets {
	return &Secrets{
		typeSecrets:     make([]models.SecretType, 0),
		secretLogPass:   make(map[models.SecretID]models.SecretLogPass),
		secretsText:     make(map[models.SecretID]models.SecretText),
		secretsBinary:   make(map[models.SecretID]models.SecretBinary),
		secretsBankCard: make(map[models.SecretID]models.SecretBankCard),
	}
}

func (s *Secrets) GetSecretTypes() []models.SecretType {
	s.mx.RLock()

	result := make([]models.SecretType, len(s.typeSecrets))

	for k, v := range s.typeSecrets {
		result[k] = v
	}

	s.mx.RUnlock()

	return result
}

// SetSecretTypes Прогрузка секретов в локальное хранилище
func (s *Secrets) SetSecretTypes(src []models.SecretType) {
	s.mx.RLock()

	s.typeSecrets = make([]models.SecretType, len(src))

	for k, v := range src {
		s.typeSecrets[k] = v
	}

	s.mx.RUnlock()
}

// CreateSecret Создание нового секрета.
func (s *Secrets) CreateSecret(secret models.SecretServerResponse) error {
	_, err := s.GetSecretByID(secret.ID, secret.SecretTypeID)

	if err == nil {
		return models.ErrSecretAlreadyExist
	}

	if !errors.Is(err, models.ErrSecretNotFound) {
		return err
	}

	secretIn, err := secret.ToSecret()
	if err != nil {
		return fmt.Errorf("secret.ToSecret: %w", err)
	}

	s.mx.Lock()
	defer s.mx.Unlock()

	switch secret.SecretTypeID {
	case models.SecretTypeIDLogpass:
		s.secretLogPass[secret.ID] = secretIn.(models.SecretLogPass)
	case models.SecretTypeIDText:
		s.secretsText[secret.ID] = secretIn.(models.SecretText)
	case models.SecretTypeIDBinary:
		s.secretsBinary[secret.ID] = secretIn.(models.SecretBinary)
	case models.SecretTypeIDBankCard:
		s.secretsBankCard[secret.ID] = secretIn.(models.SecretBankCard)
	default:
		return models.ErrUnckowType
	}

	return nil
}

// EditSecret Редактирование секрета.
func (s *Secrets) EditSecret(secret models.SecretServerResponse,
	secretID models.SecretID, typeID models.SecretTypeID) error {
	_, err := s.GetSecretByID(secretID, typeID)
	if err != nil {
		return err
	}

	newSecret, err := secret.ToSecret()
	if err != nil {
		return fmt.Errorf("secret.ToSecret: %w", err)
	}

	s.mx.Lock()
	defer s.mx.Unlock()

	switch typeID {
	case models.SecretTypeIDLogpass:
		s.secretLogPass[secretID] = newSecret.(models.SecretLogPass)
	case models.SecretTypeIDText:
		s.secretsText[secretID] = newSecret.(models.SecretText)
	case models.SecretTypeIDBinary:
		s.secretsBinary[secretID] = newSecret.(models.SecretBinary)
	case models.SecretTypeIDBankCard:
		s.secretsBankCard[secretID] = newSecret.(models.SecretBankCard)
	default:
		return models.ErrUnckowType
	}

	return nil
}

// DeleteSecret Удаление секрета
func (s *Secrets) DeleteSecret(secretID models.SecretID, typeID models.SecretTypeID) error {
	_, err := s.GetSecretByID(secretID, typeID)
	if err != nil {
		return err
	}

	s.mx.Lock()
	defer s.mx.Unlock()

	switch typeID {
	case models.SecretTypeIDLogpass:
		delete(s.secretLogPass, secretID)
	case models.SecretTypeIDText:
		delete(s.secretsText, secretID)
	case models.SecretTypeIDBinary:
		delete(s.secretsBinary, secretID)
	case models.SecretTypeIDBankCard:
		delete(s.secretsBankCard, secretID)
	default:
		return models.ErrUnckowType
	}

	return nil
}

// GetSecretByID Поиск секрета по его ID
func (s *Secrets) GetSecretByID(secretID models.SecretID, typeID models.SecretTypeID) (any, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	switch typeID {
	case models.SecretTypeIDLogpass:
		if secret, ok := s.secretLogPass[secretID]; ok {
			return secret, nil
		}

		return nil, models.ErrSecretNotFound

	case models.SecretTypeIDText:
		if secret, ok := s.secretsText[secretID]; ok {
			return secret, nil
		}

		return nil, models.ErrSecretNotFound

	case models.SecretTypeIDBinary:
		if secret, ok := s.secretsBinary[secretID]; ok {
			return secret, nil
		}

		return nil, models.ErrSecretNotFound

	case models.SecretTypeIDBankCard:
		if secret, ok := s.secretsBankCard[secretID]; ok {
			return secret, nil
		}

		return nil, models.ErrSecretNotFound

	default:
		return nil, models.ErrUnckowType
	}
}

// GetSecretsByTypeID Получение всех секретов по типу.
func (s *Secrets) GetSecretsByTypeID(typeID models.SecretTypeID) ([]any, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	switch typeID {
	case models.SecretTypeIDLogpass:
		result := make([]models.SecretLogPass, 0)
		for _, v := range s.secretLogPass {
			result = append(result, v)
		}

		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})

		return helper.SliceToAnySlice(result), nil

	case models.SecretTypeIDText:
		result := make([]models.SecretText, 0)
		for _, v := range s.secretsText {
			result = append(result, v)
		}

		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})

		return helper.SliceToAnySlice(result), nil

	case models.SecretTypeIDBinary:
		result := make([]models.SecretBinary, 0)
		for _, v := range s.secretsBinary {
			result = append(result, v)
		}

		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})

		return helper.SliceToAnySlice(result), nil

	case models.SecretTypeIDBankCard:
		result := make([]models.SecretBankCard, 0)
		for _, v := range s.secretsBankCard {
			result = append(result, v)
		}

		sort.Slice(result, func(i, j int) bool {
			return result[i].ID < result[j].ID
		})

		return helper.SliceToAnySlice(result), nil

	default:
		return nil, models.ErrUnckowType
	}

}

// SyncSecrets Полная перезапись всех секретов данными с сервера.
func (s *Secrets) SyncSecrets(ctx context.Context, src []models.SecretServerResponse) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	var once sync.Once

	for _, srcSecret := range src {
		secretIn, err := srcSecret.ToSecret()
		if err != nil {
			return fmt.Errorf("srcSecret.ToSecret: %w", err)
		}

		switch secret := secretIn.(type) {
		case models.SecretLogPass:
			onceFn := func() {
				s.secretLogPass = make(map[models.SecretID]models.SecretLogPass)
			}

			once.Do(onceFn)

			s.secretLogPass[secret.ID] = secret

		case models.SecretText:
			onceFn := func() {
				s.secretsText = make(map[models.SecretID]models.SecretText)
			}

			once.Do(onceFn)

			s.secretsText[secret.ID] = secret

		case models.SecretBinary:
			onceFn := func() {
				s.secretsBinary = make(map[models.SecretID]models.SecretBinary)
			}

			once.Do(onceFn)

			s.secretsBinary[secret.ID] = secret

		case models.SecretBankCard:
			onceFn := func() {
				s.secretsBankCard = make(map[models.SecretID]models.SecretBankCard)
			}

			once.Do(onceFn)

			s.secretsBankCard[secret.ID] = secret

		default:
			return models.ErrUnckowType
		}
	}

	return nil
}
