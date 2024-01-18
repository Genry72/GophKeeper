package models

import (
	"fmt"
	"time"
)

type CtxUserID struct{}

// Users Пользователи системы
type Users struct {
	Id           int64     `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

var (
	ErrUserAlreadyExist = fmt.Errorf("логин занят")
	ErrUserNotFound     = fmt.Errorf("пользователь с указанным логиными и паролем не найден")
	ErrEmptyLogPass     = fmt.Errorf("логин и пароль не должны быть пустыми")
)
