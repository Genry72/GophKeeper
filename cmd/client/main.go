package main

import (
	"context"
	memstorage "github.com/Genry72/GophKeeper/internal/client/repositories/memstorage/secrets"
	"github.com/Genry72/GophKeeper/internal/client/usecase/grpcclient"
	"github.com/Genry72/GophKeeper/internal/client/usecase/secrets"
	"github.com/Genry72/GophKeeper/internal/client/usecase/sync"
	"github.com/Genry72/GophKeeper/internal/client/usecase/tuiclient"
	"github.com/Genry72/GophKeeper/internal/client/usecase/users"
	"github.com/Genry72/GophKeeper/pkg/logger"
	"go.uber.org/zap"
)

const (
	grpcServerAddress = ":3200"
)

var (
	buildVersion string
	buildDate    string
)

func main() {

	zapLogger := logger.NewZapLogger("info", true)

	defer func() {
		_ = zapLogger.Sync()
	}()

	zapLogger.Info("build version:\t" + buildVersion)
	zapLogger.Info("build date:\t" + buildDate)

	ctxMain, cancelMain := context.WithCancel(context.Background())

	grpcClient, err := grpcclient.NewClient(grpcServerAddress, zapLogger)
	if err != nil {
		zapLogger.Fatal("grpcclient.NewClient", zap.Error(err))
	}

	localRepo := memstorage.NewSecrets()

	// синхронизация с сервером
	syncService := sync.NewSync(localRepo, grpcClient.SecretsClient, zapLogger)

	ucUser := users.NewUserUc(grpcClient.UsersClient, localRepo, syncService, zapLogger)

	ucSecrets := secrets.NewSecretUc(grpcClient.SecretsClient, localRepo, syncService, zapLogger)

	client := tuiclient.NewApp(ucUser, ucSecrets, zapLogger)

	if err := client.Run(ctxMain); err != nil {
		zapLogger.Fatal("client.Run", zap.Error(err))
	}

	cancelMain()

}
