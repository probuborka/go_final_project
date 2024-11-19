package http

import (
	"encoding/json"
	"net/http"

	"github.com/probuborka/go_final_project/pkg/logger"
)

func response(w http.ResponseWriter, v any, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	resp, err := json.Marshal(&v)
	if err != nil {
		logger.Error(err)
		return
	}
	w.Write(resp)
}
