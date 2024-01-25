package users

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
)

func (u *UserUc) Auth(ctx context.Context, username, password string) error {
	token, err := u.netClientUsers.Auth(ctx, username, password)
	if err != nil {
		return fmt.Errorf("u.netClientUsers.Auth: %w", err)
	}

	models.Token = token

	if err := u.sync.StartSync(ctx); err != nil {
		return fmt.Errorf("u.sync.StartSync: %w", err)
	}

	return nil
}
