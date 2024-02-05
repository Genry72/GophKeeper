package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
)

// FindByLogin Поиск пользователя по логину
func (r *UsersRepo) FindByLogin(ctx context.Context, login string) (*models.Users, bool, error) {
	query := "select id, username, password_hash, created_at from users where username = $1"
	user := &models.Users{}

	if err := r.conn.QueryRowxContext(ctx, query, login).StructScan(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("r.conn.QueryRow: %w", err)
	}

	return user, true, nil
}

// FindByID Поиск пользователя по id
func (r *UsersRepo) FindByID(ctx context.Context, id int64) (*models.Users, bool, error) {
	query := "select id, username, password_hash, created_at from users where id = $1"
	user := &models.Users{}

	if err := r.conn.QueryRowxContext(ctx, query, id).StructScan(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("r.conn.QueryRow: %w", err)
	}

	return user, true, nil
}
