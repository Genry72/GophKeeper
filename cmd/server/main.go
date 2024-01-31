package main

import (
	"github.com/Genry72/GophKeeper/cmd/server/config"
	"github.com/Genry72/GophKeeper/internal/server/handlers/grpcserver"
	"github.com/Genry72/GophKeeper/internal/server/repositories"
	"github.com/Genry72/GophKeeper/internal/server/usecase"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"github.com/Genry72/GophKeeper/pkg/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	buildVersion string
	buildDate    string
)

//type Server struct {
//	pb.UnimplementedServerServer
//}

func main() {
	zapLogger := logger.NewZapLogger("info", false)

	defer func() {
		_ = zapLogger.Sync()
	}()

	zapLogger.Info("build version:\t" + buildVersion)
	zapLogger.Info("build date:\t" + buildDate)

	conf, err := config.ReadConfig()
	if err != nil {
		zapLogger.Fatal("read config", zap.Error(err))
	}

	repo, err := repositories.NewPostgresRepo(conf.Dsn, zapLogger)
	if err != nil {
		zapLogger.Fatal("repositories.NewPostgresRepo", zap.Error(err))
	}

	jwtService := jwttoken.NewService(conf.Authkey, time.Hour)

	uc := usecase.NewUsecase(repo, jwtService, zapLogger)

	server := grpcserver.NewGrpcServer(uc, conf.ServerHostPort, jwtService, zapLogger)

	go func() {
		if err := server.Run(); err != nil {
			zapLogger.Fatal("server.Run", zap.Error(err))
		}
	}()

	// Graceful shutdown block
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-quit

	// Останавливаем сервер
	if err := server.Stop(); err != nil {
		zapLogger.Error("server.Stop", zap.Error(err))
	}

	// Закрываем подключения к БД
	repo.Stop()

}
