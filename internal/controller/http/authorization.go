package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/probuborka/go_final_project/internal/entity"
	"github.com/probuborka/go_final_project/pkg/logger"
)

type authorizationService interface {
	Password(ctx context.Context, authorization entity.Authorization) (string, error)
}

func (h handler) password(w http.ResponseWriter, r *http.Request) {
	var authorization entity.Authorization
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &authorization)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	token, err := h.authorization.Password(r.Context(), authorization)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusUnauthorized)
		//
		logger.Error(err)
		return
	}

	//
	response(w, entity.Token{Token: token}, http.StatusCreated)
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
