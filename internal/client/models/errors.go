package models

import "errors"

var (
	ErrLenLogPass         = errors.New("длина логина и пароля должны быть не менее 4-х символов")
	ErrUnckowType         = errors.New("неизвестный тип секрета")
	ErrSecretNotFound     = errors.New("секрет не найден")
	ErrSecretAlreadyExist = errors.New("секрет уже существует")
)
