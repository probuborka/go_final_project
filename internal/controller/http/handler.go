package http

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

	//createTask
	r.HandleFunc("POST /api/task", h.createTask)

	//getTasks
	r.HandleFunc("GET /api/tasks", h.getTasks)

	//getTask
	r.HandleFunc("GET /api/task", h.getTask)

	//getTask
	r.HandleFunc("PUT /api/task", h.changeTask)

	//doneTask
	r.HandleFunc("POST /api/task/done", h.doneTask)

	//deleteTask
	r.HandleFunc("DELETE /api/task", h.deleteTask)

	return r
}

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
