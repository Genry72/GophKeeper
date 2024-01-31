package grpcclient

import (
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"github.com/Genry72/GophKeeper/internal/client/usecase/grpcclient/interceptor"
	"github.com/Genry72/GophKeeper/internal/client/usecase/grpcclient/secrets"
	"github.com/Genry72/GophKeeper/internal/client/usecase/grpcclient/users"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcconn      *grpc.ClientConn
	UsersClient   *users.Users
	SecretsClient *secrets.Secrets
	log           *zap.Logger
}

func NewClient(grpcHostPort string, userInfo *models.UserInfo, log *zap.Logger) (*Client, error) {
	var interceptors []grpc.UnaryClientInterceptor
	// Добавление токена в хедер запроса
	interceptors = append(interceptors, interceptor.SetToken(userInfo.Token))
	// Кодирование и расшифровка тела секрета
	interceptors = append(interceptors, interceptor.EncodeSecretValue(userInfo.Password, log))
	grpcconn, err := grpc.Dial(grpcHostPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptors...))

	if err != nil {
		return nil, fmt.Errorf("grpc.Dial: %w", err)
	}

	return &Client{
		grpcconn:      grpcconn,
		UsersClient:   users.NewUser(grpcconn, log),
		SecretsClient: secrets.NewSecrets(grpcconn, log),
		log:           log,
	}, nil

}
