package users

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UsersRepo struct {
	conn *sqlx.DB
	log  *zap.Logger
}

func NewUsersRepo(conn *sqlx.DB, log *zap.Logger) *UsersRepo {
	return &UsersRepo{conn: conn, log: log}
}
