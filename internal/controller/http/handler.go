package http

import (
	"net/http"

	"github.com/probuborka/go_final_project/internal/entity"
)

type handler struct {
	task          task
	authorization authorization
}

func New(task task, authorization authorization) *handler {
	return &handler{
		task:          task,
		authorization: authorization,
	}
}

func (h handler) Init() *http.ServeMux {
	r := http.NewServeMux()

	//web
	r.Handle("/", http.FileServer(http.Dir(entity.WebDir)))

	//next date
	r.HandleFunc("GET /api/nextdate", h.getNextDate)

	//createTask
	r.HandleFunc("POST /api/task", auth(h.createTask))

	//getTasks
	r.HandleFunc("GET /api/tasks", auth(h.getTasks))

	//getTask
	r.HandleFunc("GET /api/task", auth(h.getTask))

	//getTask
	r.HandleFunc("PUT /api/task", auth(h.changeTask))

	//doneTask
	r.HandleFunc("POST /api/task/done", auth(h.doneTask))

	//deleteTask
	r.HandleFunc("DELETE /api/task", auth(h.deleteTask))

	//authorization
	r.HandleFunc("POST /api/signin", h.password)

	return r
}
