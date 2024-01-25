package main

import (
	"github.com/Genry72/GophKeeper/cmd/server/config"
	"github.com/Genry72/GophKeeper/internal/server/handlers/grpcserver"
	"github.com/Genry72/GophKeeper/internal/server/repositories"
	"github.com/Genry72/GophKeeper/internal/server/usecase"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"github.com/Genry72/GophKeeper/pkg/logger"
	"go.uber.org/zap"
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

	if err := server.Run(); err != nil {
		zapLogger.Fatal("server.Run", zap.Error(err))
	}

	//listen, err := net.Listen("tcp", hostPort)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// создаём gRPC-сервер, подключаем обработчики
	//s := grpc.NewServer()
	//// Регистрация серверного отражения
	//reflection.Register(s)
	//d := Server{}
	//pb.RegisterServerServer(s, &d)
	//
	//if err := s.Serve(listen); err != nil {
	//	log.Fatal(err)
	//}
}

//func (s *Server) Hello(ctx context.Context, msg *pb.HelloMsg) (*pb.HelloMsg, error) {
//	fmt.Println("Пришло на сервер " + msg.String())
//	//return &pb.HelloMsg{
//	//	Msg: "Ответ от сервера",
//	//}, nil
//	return nil, status.Errorf(codes.Aborted, "хрен тебе %s", "бугага")
//}
