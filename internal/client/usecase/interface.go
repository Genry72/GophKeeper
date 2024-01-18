package usecase

import "context"

type ITuiClient interface {
	Run() error
	Stop() error
}

// INetClient Обмен между сервером и клиентом
type INetClient interface {
	Register(ctx context.Context, username, password string) (string, error)
	Auth(ctx context.Context, username, password string) (string, error)
}
