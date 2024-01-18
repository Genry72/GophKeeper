package models

import "errors"

// UserInfo информация о пользователе
type UserInfo struct {
	Username string
	Password string
}

var (
	// ErrUserInfoEmpty нет информации о пользователе, запустившем приложение
	ErrUserInfoEmpty = errors.New("user info iz empty")
	ErrLenLogPass    = errors.New("длина логина и пароля должны быть не менее 5-ти символов")
)
