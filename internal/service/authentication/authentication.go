package authentication

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/probuborka/go_final_project/internal/entity"
)

type service struct {
}

func New() service {
	return service{}
}

func (s service) Password(ctx context.Context, authentication entity.Authentication) (string, error) {

	password := os.Getenv("TODO_PASSWORD")

	if password != authentication.Password {
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
