package grpcserver

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/server/models"
	pb "github.com/Genry72/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *GrpcServer) GetSecretTypes(ctx context.Context, in *pb.SecretTypeRequest) (*pb.SecretTypeResponse, error) {
	var response pb.SecretTypeResponse

	st, err := h.useCases.Secrets.GetSecretTypes(ctx)
	if err != nil {
		h.log.Error("h.useCases.Secrets.GetSecretTypes", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	response.SecretsType = make([]*pb.SecretType, len(st))

	for i := range st {
		response.SecretsType[i] = &pb.SecretType{
			Id:   int64(st[i].SecretTypeID),
			Name: string(st[i].SecretTypeName),
		}
	}

	return &response, nil
}

// CreateSecret Создание секрета.
func (h *GrpcServer) CreateSecret(ctx context.Context, in *pb.CreateSecretRequest) (*pb.CreateSecretResponse, error) {
	resultSecret, err := h.useCases.Secrets.AddSecret(ctx, models.SecretTypeID(in.SecretType), in.Name, in.Data)
	if err != nil {
		h.log.Error("h.useCases.Secrets.GetSecretTypes", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	response := &pb.CreateSecretResponse{
		Id:         resultSecret.ID,
		UserID:     resultSecret.UserID,
		SecretType: resultSecret.SecretTypeID,
		Name:       resultSecret.SecretName,
		Content:    resultSecret.SecretValue,
		CreatedAt:  timestamppb.New(resultSecret.CreatedAt),
		UpdatedAt:  timestamppb.New(resultSecret.UpdatedAt),
	}

	return response, nil
}

// EditSecret Редактирование секрета
func (h *GrpcServer) EditSecret(ctx context.Context, in *pb.EditSecretRequest) (*pb.EditSecretResponse, error) {
	resultSecret, err := h.useCases.Secrets.EditSecret(ctx, in.Name, in.Id, in.Data)
	if err != nil {
		h.log.Error("h.useCases.Secrets.GetSecretTypes", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	return &pb.EditSecretResponse{
		Id:         resultSecret.ID,
		UserID:     resultSecret.UserID,
		SecretType: resultSecret.SecretTypeID,
		Name:       resultSecret.SecretName,
		Content:    resultSecret.SecretValue,
		CreatedAt:  timestamppb.New(resultSecret.CreatedAt),
		UpdatedAt:  timestamppb.New(resultSecret.UpdatedAt),
	}, nil
}

func (h *GrpcServer) DeleteSecret(ctx context.Context, in *pb.DeleteSecretRequest) (*pb.DeleteSecretResponse, error) {
	if err := h.useCases.Secrets.DeleteSecret(ctx, in.Id); err != nil {
		h.log.Error("h.useCases.Secrets.GetSecretsBySecretTypeID", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	return &pb.DeleteSecretResponse{}, nil
}

func (h *GrpcServer) GetSecretsByType(ctx context.Context,
	in *pb.SecretsByTypeRequest) (*pb.SecretsByTypeResponse, error) {
	secrets, err := h.useCases.Secrets.GetSecretsBySecretTypeID(ctx, models.SecretTypeID(in.GetSecretType()))
	if err != nil {
		h.log.Error("h.useCases.Secrets.GetSecretsBySecretTypeID", zap.Error(err))
		return nil, status.Error(checkErr(err), err.Error())
	}

	resultSecrets := make([]*pb.SecretsList, len(secrets))

	for i := range secrets {
		resultSecrets[i] = &pb.SecretsList{
			Id:         secrets[i].ID,
			UserID:     secrets[i].UserID,
			SecretType: secrets[i].SecretTypeID,
			Name:       secrets[i].SecretName,
			Content:    secrets[i].SecretValue,
			CreatedAt:  timestamppb.New(secrets[i].CreatedAt),
			UpdatedAt:  timestamppb.New(secrets[i].UpdatedAt),
		}
	}

	return &pb.SecretsByTypeResponse{
		Secrets: resultSecrets,
	}, nil
}
