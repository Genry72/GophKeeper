package users

import (
	"github.com/Genry72/GophKeeper/internal/server/repositories"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"go.uber.org/zap"
)

type UsersUsecase struct {
	repo       *repositories.Repo
	jwtService *jwttoken.Service
	log        *zap.Logger
}

func NewUsersUsecase(repo *repositories.Repo, jwtService *jwttoken.Service, log *zap.Logger) *UsersUsecase {
	return &UsersUsecase{
		repo:       repo,
		jwtService: jwtService,
		log:        log,
	}
}
