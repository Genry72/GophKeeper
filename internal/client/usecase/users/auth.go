package users

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func (u *UserUc) Auth(ctx context.Context, username, password string) error {
	token, err := u.netClientUsers.Auth(ctx, username, password)
	if err != nil {
		return fmt.Errorf("u.netClientUsers.Auth: %w", err)
	}
	*u.UserInfo.Username = username
	*u.UserInfo.Password = password
	*u.UserInfo.Token = token

	if err := u.sync.StartSync(ctx); err != nil {
		return fmt.Errorf("u.sync.StartSync: %w", err)
	}

	// Запуск обновления токена
	u.updateTokenStart(ctx)

	return nil
}

// updateTokenStart Периодическое обновление токена
func (u *UserUc) updateTokenStart(ctx context.Context) {
	u.log.Info("Start periodic update token")
	// Время жизни токена 1 час, раз в 45 минут получаем новый
	t := time.NewTicker(45 * time.Minute)

	go func() {
		u.updateTokenStarted = true

		for {
			select {
			case <-t.C:
				if err := u.Auth(ctx, *u.UserInfo.Username, *u.UserInfo.Password); err != nil {
					u.log.Error("Periodic update token", zap.Error(err))
				}
			case <-ctx.Done():
				u.doneUpdateToken <- struct{}{}
				close(u.doneUpdateToken)
				return
			}
		}
	}()
}

// UpdateTokenStop Остановка обновления токена.
func (u *UserUc) UpdateTokenStop(t time.Duration) {
	u.log.Info("Stopping update token")

	if !u.updateTokenStarted {
		u.log.Info("UpdateToken not started")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), t)

	defer cancel()

	select {
	case <-ctx.Done():
		u.log.Error("Update token not stopped")
	case <-u.doneUpdateToken:
		u.updateTokenStarted = false
		u.log.Info("Update token success stopped")
	}
}
