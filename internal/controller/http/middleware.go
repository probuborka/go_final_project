package http

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
	"github.com/probuborka/go_final_project/pkg/logger"
)

type middleware func(http.Handler) http.Handler

func compileMiddleware(h http.Handler, m []middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger.Infof("%s %s %s", req.Method, req.RequestURI, time.Now())
		next.ServeHTTP(w, req)
	})
}

func authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// смотрим наличие пароля
		if len(cfg.Password) > 0 {
			var token string // JWT-токен из куки
			// получаем куку
			cookie, err := r.Cookie("token")
			if err == nil {
				token = cookie.Value
			}
			//
			secret := []byte(cfg.Password)
			// здесь код для валидации и проверки JWT-токена
			jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
				// секретный ключ для всех токенов одинаковый, поэтому просто возвращаем его
				return secret, nil
			})
			if err != nil {
				response(w, entityerror.Error{Error: err.Error()}, http.StatusUnauthorized)
				return
			}
			if !jwtToken.Valid {
				// возвращаем ошибку авторизации 401
				response(w, entityerror.Error{Error: "Authentification required"}, http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// func auth(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// смотрим наличие пароля
// 		pass := os.Getenv("TODO_PASSWORD")
// 		if len(pass) > 0 {
// 			var token string // JWT-токен из куки
// 			// получаем куку
// 			cookie, err := r.Cookie("token")
// 			if err == nil {
// 				token = cookie.Value
// 			}
// 			//
// 			secret := []byte(pass)
// 			// здесь код для валидации и проверки JWT-токена
// 			jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 				// секретный ключ для всех токенов одинаковый, поэтому просто возвращаем его
// 				return secret, nil
// 			})
// 			if err != nil {
// 				response(w, entity.Error{Error: err.Error()}, http.StatusUnauthorized)
// 				return
// 			}
// 			if !jwtToken.Valid {
// 				// возвращаем ошибку авторизации 401
// 				response(w, entity.Error{Error: "Authentification required"}, http.StatusUnauthorized)
// 				return
// 			}
// 		}
// 		next(w, r)
// 	})
// }
