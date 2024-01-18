package grpcserver

import (
	"errors"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/handlers/grpcserver/interceptor"
	"github.com/Genry72/GophKeeper/internal/server/models"
	"github.com/Genry72/GophKeeper/internal/server/usecase"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"github.com/Genry72/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"net"
)

type GrpcServer struct {
	useCases   *usecase.Usecase
	grpcServer *grpc.Server
	log        *zap.Logger
	hostPort   string
	listener   net.Listener // Сетевое соединение
	proto.UnimplementedUsersServer
}

func NewGrpcServer(useCases *usecase.Usecase,
	hostPort string, jwtService *jwttoken.Service, log *zap.Logger) *GrpcServer {
	// создаём gRPC-сервер, подключаем обработчики
	interceptors := make([]grpc.UnaryServerInterceptor, 0)

	// Логирование
	interceptors = append(interceptors, interceptor.Logging(log))

	// Проверка заголовка с токеном
	interceptors = append(interceptors, interceptor.CheckToken(jwtService, log))

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors...))

	// Регистрация серверного отражения
	reflection.Register(server)

	grpStruct := &GrpcServer{
		useCases:   useCases,
		grpcServer: server,
		hostPort:   hostPort,
		log:        log,
	}

	// регистрируем сервис
	proto.RegisterUsersServer(server, grpStruct)

	return grpStruct
}

func (s *GrpcServer) Run() error {
	listen, err := net.Listen("tcp", s.hostPort)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	s.listener = listen

	// получаем запрос gRPC
	if err := s.grpcServer.Serve(listen); err != nil {
		return fmt.Errorf("s.grpcServer.Serve: %w", err)
	}

	return nil
}

func (s *GrpcServer) Stop() error {
	s.grpcServer.GracefulStop()
	if err := s.listener.Close(); err != nil {
		return fmt.Errorf("s.listener.Close: %w", err)
	}

	s.log.Info("server success stopped")
	return nil
}

func checkErr(err error) codes.Code {
	switch {
	case errors.Is(err, models.ErrUserAlreadyExist):
		return codes.AlreadyExists
	case errors.Is(err, models.ErrUserNotFound):
		return codes.Unauthenticated
	default:
		return codes.Internal
	}
}
