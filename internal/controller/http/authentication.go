package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/probuborka/go_final_project/internal/entity"
	"github.com/probuborka/go_final_project/pkg/logger"
)

type serviceAuthentication interface {
	Password(ctx context.Context, authentication entity.Authentication) (string, error)
}

func (h handler) password(w http.ResponseWriter, r *http.Request) {
	var authentication entity.Authentication
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &authentication)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	token, err := h.authentication.Password(r.Context(), authentication)
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
