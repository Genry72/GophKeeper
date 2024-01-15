package jwttoken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int64
}

type Service struct {
	tokenKey string
	lifetime time.Duration
}

func NewService(tokenKey string, lifetime time.Duration) *Service {
	return &Service{tokenKey: tokenKey, lifetime: lifetime}
}

// GetToken Получение токена по id пользователя
func (j *Service) GetToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.lifetime)),
		},

		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(j.tokenKey))
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}

	return tokenString, nil
}

// ValidateAndParseToken Проверка корректности токена. Возвращается id пользователя
func (j *Service) ValidateAndParseToken(token string) (int64, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.tokenKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("jwt.ParseWithClaims: %w", err)
	}

	return claims.UserID, nil
}
