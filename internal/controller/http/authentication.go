package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	entityauth "github.com/probuborka/go_final_project/internal/entity/authentication"
	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
	"github.com/probuborka/go_final_project/pkg/logger"
)

type serviceAuthentication interface {
	Password(ctx context.Context, authentication entityauth.Authentication) (string, error)
}

func (h handler) password(w http.ResponseWriter, r *http.Request) {
	var authentication entityauth.Authentication
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &authentication)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	token, err := h.authentication.Password(r.Context(), authentication)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusUnauthorized)
		//
		logger.Error(err)
		return
	}

	//
	response(w, entityauth.Token{Token: token}, http.StatusCreated)
}
