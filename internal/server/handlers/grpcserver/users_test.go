package grpcserver

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/server/models"
	"github.com/Genry72/GophKeeper/internal/server/repositories"
	mockrepo "github.com/Genry72/GophKeeper/internal/server/repositories/mocks"
	"github.com/Genry72/GophKeeper/internal/server/usecase"
	"github.com/Genry72/GophKeeper/internal/server/usecase/secrets"
	"github.com/Genry72/GophKeeper/internal/server/usecase/users"
	"github.com/Genry72/GophKeeper/pkg/hash"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"github.com/Genry72/GophKeeper/pkg/logger"
	pb "github.com/Genry72/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"testing"
	"time"
)

func TestGrpcMethods(t *testing.T) {
	var (
		username      = "username"
		secretName    = "secret1"
		userID        = int64(1)
		pass          = "pass"
		passHashed, _ = hash.Sha512(pass)
	)

	const bufSize = 1024 * 1024

	lis := bufconn.Listen(bufSize)

	mockCtl := gomock.NewController(t)

	defer mockCtl.Finish()

	mockUsers := mockrepo.NewMockIUsers(mockCtl)
	mockSecrets := mockrepo.NewMockISecrets(mockCtl)

	log := logger.NewZapLogger("info", false)

	repo := &repositories.Repo{
		Users:   mockUsers,
		Secrets: mockSecrets,
	}

	jwtService := jwttoken.NewService("authKey", time.Hour)
	userToken, err := jwtService.GetToken(userID)
	if err != nil {
		log.Fatal(err.Error())
	}
	uu := users.NewUsersUsecase(repo, jwtService, log)
	us := secrets.NewSecretsUsecase(repo, log)

	server := NewGrpcServer(&usecase.Usecase{
		Users:   uu,
		Secrets: us,
	}, "bufnet", jwtService, log)

	ctx, cancelCtc := context.WithTimeout(context.Background(), time.Second)
	ctx = metadata.AppendToOutgoingContext(ctx, models.HeaderAuthorization, fmt.Sprintf("Bearer %s", userToken))
	defer cancelCtc()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	go func() {
		if err := server.grpcServer.Serve(lis); err != nil {
			log.Fatal(err.Error())
		}
	}()

	clientUsers := pb.NewUsersClient(conn)
	clientSecrets := pb.NewSecretClient(conn)

	tests := []struct {
		name   string
		testFn func()
		mockFn func()
	}{
		{
			name: "register user",
			mockFn: func() {
				mockUsers.EXPECT().Register(gomock.Any(), username, gomock.Any()).Return(int64(123), nil)
				mockUsers.EXPECT().FindByLogin(gomock.Any(), username).Return(&models.Users{PasswordHash: passHashed}, true, nil)
			},
			testFn: func() {
				resp, err := clientUsers.Register(ctx, &pb.RegisterUserMsg{Username: username, Password: pass})
				if err != nil {
					assert.NoError(t, err)
				}
				assert.NotEqual(t, "", resp.Token)
			},
		},
		{
			name: "auth user",
			mockFn: func() {
				mockUsers.EXPECT().FindByLogin(gomock.Any(), username).Return(&models.Users{PasswordHash: passHashed}, true, nil)
			},
			testFn: func() {
				resp, err := clientUsers.Auth(ctx, &pb.AuthUserMsg{Username: username, Password: pass})
				if err != nil {
					assert.NoError(t, err)
				}
				assert.NotEqual(t, "", resp.Token)
			},
		},
		{
			name: "secrets by typeID",
			mockFn: func() {
				mockSecrets.EXPECT().GetSecretsBySecretTypeID(gomock.Any(), userID, int64(1)).Return([]models.Secret{
					{

						SecretTypeID: 1,
					}}, nil)
			},
			testFn: func() {
				resp, err := clientSecrets.GetSecretsByType(ctx, &pb.SecretsByTypeRequest{
					SecretType: 1,
				})

				if err != nil {
					assert.NoError(t, err)
				}
				assert.Equal(t, int64(1), resp.Secrets[0].SecretType)
			},
		},
		{
			name: "GetSecretTypes",
			mockFn: func() {
				mockSecrets.EXPECT().GetSecretTypes(gomock.Any()).Return([]models.SecretType{
					{
						SecretTypeID:   1,
						SecretTypeName: "name1",
					},
					{
						SecretTypeID:   2,
						SecretTypeName: "name2",
					},
				}, nil)
			},
			testFn: func() {
				resp, err := clientSecrets.GetSecretTypes(ctx, &pb.SecretTypeRequest{})

				if err != nil {
					assert.NoError(t, err)
				}
				assert.Equal(t, []*pb.SecretType{
					{
						Id:   1,
						Name: "name1",
					},
					{
						Id:   2,
						Name: "name2",
					},
				}, resp.SecretsType)
			},
		},
		{
			name: "CreateSecret",
			mockFn: func() {
				mockSecrets.EXPECT().AddSecret(gomock.Any(), userID, int64(1), "secret1", []byte("11")).Return(models.Secret{
					ID:           1,
					UserID:       userID,
					SecretTypeID: 1,
					SecretName:   "secret1",
					SecretValue:  []byte("11"),
					CreatedAt: time.Date(2024, 01, 04, 22, 00, 00, 0,
						time.Local),
					UpdatedAt: time.Date(2024, 01, 04, 22, 00, 00, 0,
						time.Local),
					DeletedAt: nil,
				}, nil)
			},
			testFn: func() {
				resp, err := clientSecrets.CreateSecret(ctx, &pb.CreateSecretRequest{
					Name:       "secret1",
					SecretType: 1,
					Data:       []byte("11"),
				})

				if err != nil {
					assert.NoError(t, err)
				}
				want := &pb.CreateSecretResponse{
					Id:         1,
					UserID:     userID,
					SecretType: 1,
					Name:       "secret1",
					Content:    []byte("11"),
					CreatedAt: timestamppb.New(time.Date(2024, 01, 04, 22, 00, 00,
						0, time.Local)),
					UpdatedAt: timestamppb.New(time.Date(2024, 01, 04, 22, 00, 00,
						0, time.Local)),
				}
				assert.Equal(t, want.Id, resp.Id)
				assert.Equal(t, want.UserID, resp.UserID)
				assert.Equal(t, want.SecretType, resp.SecretType)
				assert.Equal(t, want.Name, resp.Name)
				assert.Equal(t, want.Content, resp.Content)
				assert.Equal(t, want.CreatedAt, resp.CreatedAt)
				assert.Equal(t, want.UpdatedAt, resp.UpdatedAt)

			},
		},
		{
			name: "EditSecret",
			mockFn: func() {
				mockSecrets.EXPECT().GetSecretByID(gomock.Any(), userID, int64(1)).Return(models.Secret{
					ID:           1,
					UserID:       userID,
					SecretTypeID: 1,
					SecretName:   "secret1",
					SecretValue:  []byte("11"),
					CreatedAt: time.Date(2024, 01, 04, 22, 00, 00, 0,
						time.Local),
					UpdatedAt: time.Date(2024, 01, 04, 22, 00, 00, 0,
						time.Local),
					DeletedAt: nil,
				}, nil)
				mockSecrets.EXPECT().EditSecret(gomock.Any(), secretName, int64(1), []byte("11")).Return(models.Secret{
					ID:           1,
					UserID:       userID,
					SecretTypeID: 1,
					SecretName:   "secret1",
					SecretValue:  []byte("11"),
					CreatedAt: time.Date(2024, 01, 04, 22, 00, 00, 0,
						time.Local),
					UpdatedAt: time.Date(2024, 01, 04, 22, 00, 00, 0,
						time.Local),
					DeletedAt: nil,
				}, nil)
			},
			testFn: func() {
				resp, err := clientSecrets.EditSecret(ctx, &pb.EditSecretRequest{
					Id:   int64(1),
					Name: "secret1",
					Data: []byte("11"),
				})

				if err != nil {
					assert.NoError(t, err)
				}
				want := &pb.CreateSecretResponse{
					Id:         1,
					UserID:     userID,
					SecretType: 1,
					Name:       "secret1",
					Content:    []byte("11"),
					CreatedAt: timestamppb.New(time.Date(2024, 01, 04, 22, 00, 00,
						0, time.Local)),
					UpdatedAt: timestamppb.New(time.Date(2024, 01, 04, 22, 00, 00,
						0, time.Local)),
				}
				assert.Equal(t, want.Id, resp.Id)
				assert.Equal(t, want.UserID, resp.UserID)
				assert.Equal(t, want.SecretType, resp.SecretType)
				assert.Equal(t, want.Name, resp.Name)
				assert.Equal(t, want.Content, resp.Content)
				assert.Equal(t, want.CreatedAt, resp.CreatedAt)
				assert.Equal(t, want.UpdatedAt, resp.UpdatedAt)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn()
			}
			tt.testFn()
		})
	}
}
