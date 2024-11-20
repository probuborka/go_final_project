package service

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/probuborka/go_final_project/internal/entity"
)

type authorization struct {
}

func newAuthorization() authorization {
	return authorization{}
}

func (a authorization) Password(ctx context.Context, authorization entity.Authorization) (string, error) {

	password := os.Getenv("TODO_PASSWORD")

	if password != authorization.Password {
		return "", entity.ErrInvalidPassword
	}

	secret := []byte(password)

	// создаём jwt и указываем алгоритм хеширования
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	// получаем подписанный токен
	token, err := jwtToken.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
