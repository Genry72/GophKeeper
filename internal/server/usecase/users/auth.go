package users

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
	"github.com/Genry72/GophKeeper/pkg/hash"
)

// AuthUser Аутентификация пользователя. Возвращает токен
func (u *UsersUsecase) AuthUser(ctx context.Context, login, pass string) (string, error) {
	userFromDB, ok, err := u.repo.Users.FindByLogin(ctx, login)
	if err != nil {
		return "", fmt.Errorf("u.repo.Users.FindByLogin: %w", err)
	}

	if !ok {
		return "", models.ErrUserNotFound
	}

	passHashed, err := hash.Sha512(pass)
	if err != nil {
		return "", fmt.Errorf("hash.Sha512: %w", err)
	}

	if userFromDB.PasswordHash != passHashed {
		return "", models.ErrUserNotFound
	}

	return u.jwtService.GetToken(userFromDB.Id)
}
