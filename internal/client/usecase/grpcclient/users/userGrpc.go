package users

import (
	"github.com/Genry72/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Users struct {
	usersClient proto.UsersClient
	log         *zap.Logger
}

func NewUser(grpcconn grpc.ClientConnInterface, log *zap.Logger) *Users {
	usersClient := proto.NewUsersClient(grpcconn)

	return &Users{
		usersClient: usersClient,
		log:         log,
	}
}
