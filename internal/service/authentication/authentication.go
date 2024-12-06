package authentication

import (
	"context"

	"github.com/golang-jwt/jwt"
	entityauth "github.com/probuborka/go_final_project/internal/entity/authentication"
	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
)

type service struct {
	auth entityconfig.Authentication
}

func New(auth entityconfig.Authentication) service {
	return service{
		auth: auth,
	}
}

func (s service) Password(ctx context.Context, authentication entityauth.Authentication) (string, error) {

	if s.auth.Password != authentication.Password {
		return "", entityerror.ErrInvalidPassword
	}

	secret := []byte(s.auth.Password)

	// создаём jwt и указываем алгоритм хеширования
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	// получаем подписанный токен
	token, err := jwtToken.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
