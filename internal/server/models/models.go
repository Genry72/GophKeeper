package models

import (
	"fmt"
	"time"
)

var (
	ErrUserAlreadyExist = fmt.Errorf("логин занят")
	ErrUserNotFound     = fmt.Errorf("пользователь с указанным логиными и паролем не найден")
)

// Users Пользователи системы
type Users struct {
	Id           int64     `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}
