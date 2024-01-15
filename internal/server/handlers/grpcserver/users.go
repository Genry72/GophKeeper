package grpcserver

import (
	"context"
	pb "github.com/Genry72/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

// Register Регистрация пользователя
func (h *GrpcServer) Register(ctx context.Context, in *pb.RegisterUserMsg) (*pb.TokenResponse, error) {
	var response pb.TokenResponse

	if _, err := h.useCases.Users.RegisterUser(ctx, in.Username, in.Password); err != nil {
		h.log.Error("h.useCases.Users.RegisterUser", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	token, err := h.useCases.Users.AuthUser(ctx, in.Username, in.Password)
	if err != nil {
		h.log.Error("h.useCases.Users.AuthUser", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	response.Token = token

	return &response, nil
}

// Auth Аутентификация пользователя
func (h *GrpcServer) Auth(ctx context.Context, in *pb.AuthUserMsg) (*pb.TokenResponse, error) {
	var response pb.TokenResponse

	token, err := h.useCases.Users.AuthUser(ctx, in.Username, in.Password)
	if err != nil {
		h.log.Error("h.useCases.Users.AuthUser", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	response.Token = token

	return &response, nil
}
