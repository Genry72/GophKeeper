package grpcclient

import (
	"fmt"
	"github.com/Genry72/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcconn    *grpc.ClientConn
	usersClient proto.UsersClient
	log         *zap.Logger
}

func NewClient(grpcHostPort string, log *zap.Logger) (*Client, error) {

	//grpcconn, err := grpc.Dial(grpcHostPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
	//	grpc.WithChainUnaryInterceptor(interceptors...))

	grpcconn, err := grpc.Dial(grpcHostPort, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor())

	if err != nil {
		return nil, fmt.Errorf("grpc.Dial: %w", zap.Error(err))
	}

	usersClient := proto.NewUsersClient(grpcconn)

	return &Client{
		grpcconn:    grpcconn,
		usersClient: usersClient,
		log:         log,
	}, nil

}
