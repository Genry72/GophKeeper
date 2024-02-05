package users

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/pkg/hash"
)

// RegisterUser Регистрация пользователя
func (u *UsersUsecase) RegisterUser(ctx context.Context, login, pass string) (int64, error) {
	passHashed, err := hash.Sha512(pass)
	if err != nil {
		return 0, fmt.Errorf("hash.Sha512: %w", err)
	}
	return u.repo.Users.Register(ctx, login, passHashed)
}
