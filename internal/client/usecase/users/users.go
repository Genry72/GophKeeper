package users

import (
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/internal/client/repositories"
	"github.com/Genry72/GophKeeper/internal/client/usecase"
	"go.uber.org/zap"
)

type UserUc struct {
	netClientUsers usecase.INetClientUsers   // Сетевой клиент для обмена с севрером
	localRepo      repositories.IrepoSecrets // Локальное хранение секретов
	sync           usecase.ISync             // Синхронизация данных с сервером
	UserInfo       *models.UserInfo          // Информация о пользователе, запустившем приложение
	log            *zap.Logger
}

func NewUserUc(userInfo *models.UserInfo, netClient usecase.INetClientUsers, localRepo repositories.IrepoSecrets,
	sync usecase.ISync, log *zap.Logger) *UserUc {

	return &UserUc{
		netClientUsers: netClient,
		localRepo:      localRepo,
		sync:           sync,
		UserInfo:       userInfo,
		log:            log,
	}
}
