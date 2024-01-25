package users

import (
	"github.com/Genry72/GophKeeper/internal/client/repositories"
	"github.com/Genry72/GophKeeper/internal/client/usecase"
	"go.uber.org/zap"
)

type UserUc struct {
	netClientUsers usecase.INetClientUsers   // Сетевой клиент для обмена с севрером
	localRepo      repositories.IrepoSecrets // Локальное хранение секретов
	sync           usecase.ISync             // Синхронизация данных с сервером
	log            *zap.Logger
}

func NewUserUc(netClient usecase.INetClientUsers, localRepo repositories.IrepoSecrets,
	sync usecase.ISync, log *zap.Logger) *UserUc {
	return &UserUc{
		netClientUsers: netClient,
		localRepo:      localRepo,
		sync:           sync,
		log:            log,
	}
}
