package interceptor

import (
	"context"
	"fmt"
	"github.com/Genry72/GophKeeper/internal/client/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SetToken Добавление токена во все исходящие запросы
func SetToken(token *string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{},
		reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		if token != nil {
			ctx = metadata.AppendToOutgoingContext(ctx, models.HeaderAuthorization, fmt.Sprintf("Bearer %s", *token))
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
