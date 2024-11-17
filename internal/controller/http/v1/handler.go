package v1

import (
	"encoding/json"
	"net/http"

	"github.com/probuborka/go_final_project/internal/entity"
	"github.com/probuborka/go_final_project/pkg/logger"
)

type handler struct {
	task task
}

func New(task task) *handler {
	return &handler{
		task: task,
	}
}

func (h handler) Init() *http.ServeMux {
	r := http.NewServeMux()

	//web
	r.Handle("/", http.FileServer(http.Dir(entity.WebDir)))

	//next date
	r.HandleFunc("GET /api/nextdate", h.getNextDate)

	//task
	r.HandleFunc("POST /api/task", h.createTask)

	return r
}

func response(w http.ResponseWriter, v any, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		logger.Error(err)
	}
}
