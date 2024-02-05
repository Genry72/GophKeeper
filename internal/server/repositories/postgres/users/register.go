package users

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
)

// Register Проверка на существование логина и запись в базу в случае его отсутствия.
// Возвращается id созданного пользователя
func (r *UsersRepo) Register(ctx context.Context, login, encryptPass string) (int64, error) {
	_, ok, err := r.FindByLogin(ctx, login)
	if err != nil {
		return 0, fmt.Errorf("r.FindByLogin: %w", err)
	}

	if ok {
		return 0, models.ErrUserAlreadyExist
	}

	var id int64

	query := "INSERT INTO users (username, password_hash) VALUES ($1, $2) returning id"

	if err := r.conn.QueryRow(query, login, encryptPass).Scan(&id); err != nil {
		r.log.Error(query)
		return 0, fmt.Errorf("r.conn.QueryRow: %w", err)
	}

	return id, nil
}
