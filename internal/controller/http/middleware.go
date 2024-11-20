package http

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/probuborka/go_final_project/internal/entity"
	"github.com/probuborka/go_final_project/pkg/logger"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		logger.Infof("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// смотрим наличие пароля
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var token string // JWT-токен из куки
			// получаем куку
			cookie, err := r.Cookie("token")
			if err == nil {
				token = cookie.Value
			}
			//
			secret := []byte(pass)
			// здесь код для валидации и проверки JWT-токена
			jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
				// секретный ключ для всех токенов одинаковый, поэтому просто возвращаем его
				return secret, nil
			})
			if err != nil {
				response(w, entity.Error{Error: err.Error()}, http.StatusUnauthorized)
				return
			}
			if !jwtToken.Valid {
				// возвращаем ошибку авторизации 401
				response(w, entity.Error{Error: "Authentification required"}, http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
