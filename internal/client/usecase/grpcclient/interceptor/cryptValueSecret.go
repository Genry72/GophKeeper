package interceptor

import (
	"context"
	cryptor "github.com/Genry72/GophKeeper/pkg/crypt"
	"github.com/Genry72/GophKeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// EncodeSecretValue Шифрование данных тела секрета.
func EncodeSecretValue(password *string, log *zap.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{},
		reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		// Кодируем тело секрета в исходящем запросе
		switch decryptedRequest := req.(type) {
		case *proto.CreateSecretRequest:
			encrypdedBody, err := cryptor.Encrypt(decryptedRequest.Data, *password)
			if err != nil {
				log.Error("cryptor.Encrypt", zap.Error(err))
				return err
			}

			decryptedRequest.Data = []byte(encrypdedBody)
			req = decryptedRequest
		case *proto.EditSecretRequest:
			encrypdedBody, err := cryptor.Encrypt(decryptedRequest.Data, *password)
			if err != nil {
				log.Error("cryptor.Encrypt", zap.Error(err))
				return err
			}

			decryptedRequest.Data = []byte(encrypdedBody)
			req = decryptedRequest
		}

		// Вызываем метод
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			log.Error("invoker", zap.Error(err))
			return err
		}

		// Расшифровываем тело вернувшегося секрета при ответе

		switch encryptedResponse := reply.(type) {
		case *proto.CreateSecretResponse:
			decryptedBody, err := cryptor.Decrypt(string(encryptedResponse.Content), *password)
			if err != nil {
				log.Error("cryptor.Decrypt", zap.Error(err))
				return err
			}

			encryptedResponse.Content = decryptedBody

		case *proto.EditSecretResponse:
			decryptedBody, err := cryptor.Decrypt(string(encryptedResponse.Content), *password)
			if err != nil {
				log.Error("cryptor.Decrypt", zap.Error(err))
				return err
			}

			encryptedResponse.Content = decryptedBody

		case *proto.SecretsByTypeResponse:
			for i := range encryptedResponse.Secrets {
				decryptedBody, err := cryptor.Decrypt(string(encryptedResponse.Secrets[i].Content), *password)
				if err != nil {
					log.Error("cryptor.Decrypt", zap.Error(err))
					return err
				}

				encryptedResponse.Secrets[i].Content = decryptedBody
			}
		}

		return err
	}
}
