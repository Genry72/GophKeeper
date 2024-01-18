package interceptor

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/server/models"
	"github.com/Genry72/GophKeeper/pkg/jwttoken"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

// CheckToken Валидация токена и добавление id пользователя в контекст запроса
func CheckToken(jwtService *jwttoken.Service, log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		// Не вызываем проверку для методов авторизации и аутентификации
		if info.FullMethod == "/proto.Users/Auth" || info.FullMethod == "/proto.Users/Register" {
			return handler(ctx, req)
		}

		var token string

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			values := md.Get("Authorization")
			if len(values) > 0 {
				tokenFromHeader := values[0]
				tokenStrs := strings.Split(tokenFromHeader, "Bearer ")
				if len(tokenStrs) != 2 {
					log.Error("bad auth header", zap.String("header", tokenFromHeader))
					return nil, status.Errorf(codes.Unauthenticated, "bad Authorization header")
				}
				token = tokenStrs[1]
			}
		}

		userID, err := jwtService.ValidateAndParseToken(token)
		if err != nil {
			log.Error("jwtService.ValidateAndParseToken", zap.Error(err))
			return nil, status.Errorf(codes.Unauthenticated, "user not authenticate")
		}

		ctx = context.WithValue(ctx, models.CtxUserID{}, userID)

		return handler(ctx, req)
	}
}
