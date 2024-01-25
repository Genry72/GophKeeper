package secrets

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type SecretsRepo struct {
	conn *sqlx.DB
	log  *zap.Logger
}

func NewSecretsRepo(conn *sqlx.DB, log *zap.Logger) *SecretsRepo {
	return &SecretsRepo{conn: conn, log: log}
}
