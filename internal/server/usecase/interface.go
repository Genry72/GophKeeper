package usecase

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/server/repositories"
	"github.com/Genry72/GophKeeper/internal/server/usecase/users"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"go.uber.org/zap"
)

// Iusers Работа с пользователями
type Iusers interface {
	RegisterUser(ctx context.Context, login, pass string) (int64, error)
	AuthUser(ctx context.Context, login, pass string) (string, error)
}

type Usecase struct {
	Users Iusers
	log   *zap.Logger
}

func NewUsecase(repo *repositories.Repo, jwtService *jwttoken.Service, log *zap.Logger) *Usecase {
	return &Usecase{
		Users: users.NewUsersUsecase(repo, jwtService, log),
		log:   log,
	}
}
