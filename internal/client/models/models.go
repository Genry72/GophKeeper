package models

import (
	"encoding/json"
	"fmt"
	"time"
)

const HeaderAuthorization = "Authorization"

const (
	SecretTypeIDLogpass  SecretTypeID = 1
	SecretTypeIDText     SecretTypeID = 2
	SecretTypeIDBinary   SecretTypeID = 3
	SecretTypeIDBankCard SecretTypeID = 4
)

// UserInfo информация о пользователе
type UserInfo struct {
	Username string
	Password string
}

// Token Хранение и обновление токена.
var Token = ""

// SecretServerResponse структура по обмену секретов с сервером.
type SecretServerResponse struct {
	ID           SecretID
	SecretTypeID SecretTypeID
	Name         SecretName
	Value        []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ToSecret Преобразование из серверной структуры в клиентскую
func (s SecretServerResponse) ToSecret() (any, error) {
	switch s.SecretTypeID {
	case SecretTypeIDLogpass:
		var val SecretLogPassValue

		if err := json.Unmarshal(s.Value, &val); err != nil {
			return nil, fmt.Errorf("json.Unmarshal SecretLogPassValue: %w", err)
		}

		result := SecretLogPass{
			ID:           s.ID,
			SecretTypeID: s.SecretTypeID,
			Name:         s.Name,
			Value:        val,
			CreatedAt:    s.CreatedAt,
			UpdatedAt:    s.UpdatedAt,
		}

		return result, nil

	case SecretTypeIDText:
		var val SecretTextValue

		if err := json.Unmarshal(s.Value, &val); err != nil {
			return nil, fmt.Errorf("json.Unmarshal SecretTextValue: %w", err)
		}

		result := SecretText{
			ID:           s.ID,
			SecretTypeID: s.SecretTypeID,
			Name:         s.Name,
			Value:        val,
			CreatedAt:    s.CreatedAt,
			UpdatedAt:    s.UpdatedAt,
		}

		return result, nil

	case SecretTypeIDBinary:
		result := SecretBinary{
			ID:           s.ID,
			SecretTypeID: s.SecretTypeID,
			Name:         s.Name,
			Value:        s.Value,
			CreatedAt:    s.CreatedAt,
			UpdatedAt:    s.UpdatedAt,
		}

		return result, nil

	case SecretTypeIDBankCard:
		var val SecretBankCardValue

		if err := json.Unmarshal(s.Value, &val); err != nil {
			return nil, fmt.Errorf("json.Unmarshal SecretBankCardValue: %w", err)
		}

		result := SecretBankCard{
			ID:           s.ID,
			SecretTypeID: s.SecretTypeID,
			Name:         s.Name,
			Value:        val,
			CreatedAt:    s.CreatedAt,
			UpdatedAt:    s.UpdatedAt,
		}

		return result, nil
	default:
		return nil, ErrUnckowType
	}
}

type SecretTypeID int64

type SecretName string

type SecretTypeName string

type SecretID int64

type SecretType struct {
	SecretTypeID   SecretTypeID
	SecretTypeName SecretTypeName
}

// SecretLogPassValue хранится значение секрета типа login/password
type SecretLogPassValue struct {
	Login    string
	Password string
}

type SecretTextValue string

type SecretBinaryValue []byte

// CardDateTo Срок действия карты
type CardDateTo struct {
	Year  int
	Month int
}

type SecretBankCardValue struct {
	CardNumber int64
	CardDateTo CardDateTo
	Cvv        int
}

// SecretLogPass секрет, хранящий логин и пароль
type SecretLogPass struct {
	ID           SecretID
	SecretTypeID SecretTypeID
	Name         SecretName
	Value        SecretLogPassValue
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// SecretText секрет, хранящий текстовое значение
type SecretText struct {
	ID           SecretID
	SecretTypeID SecretTypeID
	Name         SecretName
	Value        SecretTextValue
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// SecretBinary секрет, хранящий бинарные данные
type SecretBinary struct {
	ID           SecretID
	SecretTypeID SecretTypeID
	Name         SecretName
	Value        SecretBinaryValue
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// SecretBankCard секрет, хранящий банковские карты
type SecretBankCard struct {
	ID           SecretID
	SecretTypeID SecretTypeID
	Name         SecretName
	Value        SecretBankCardValue
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
