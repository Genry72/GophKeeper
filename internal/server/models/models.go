package models

import (
	"fmt"
	"time"
)

type CtxUserID struct{}

const HeaderAuthorization = "Authorization"

// Users Пользователи системы
type Users struct {
	Id           int64     `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

// Secret Секрет пользователя.
type Secret struct {
	ID           int64      `db:"id"`
	UserID       int64      `db:"user_id"`
	SecretTypeID int64      `db:"secret_type_id"`
	SecretName   string     `db:"name"`
	SecretValue  []byte     `db:"data"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}

var (
	ErrUserAlreadyExist = fmt.Errorf("логин занят")
	ErrUserNotFound     = fmt.Errorf("пользователь с указанным логиными и паролем не найден")
	ErrEmptyLogPass     = fmt.Errorf("логин и пароль не должны быть пустыми")
)

type SecretTypeID int64
type SecretName string
type SecretTypeName string

type SecretType struct {
	SecretTypeID   SecretTypeID   `db:"id"`
	SecretTypeName SecretTypeName `db:"name"`
}
