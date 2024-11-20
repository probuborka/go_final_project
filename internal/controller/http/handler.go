package http

import (
	"net/http"

	"github.com/probuborka/go_final_project/internal/entity"
)

type handler struct {
	task          taskService
	authorization authorizationService
}

func New(task taskService, authorization authorizationService) *handler {
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

	//create task
	r.HandleFunc("POST /api/task", auth(h.createTask))

	//get tasks
	r.HandleFunc("GET /api/tasks", auth(h.getTasks))

	//get task
	r.HandleFunc("GET /api/task", auth(h.getTask))

	//change task
	r.HandleFunc("PUT /api/task", auth(h.changeTask))

	//done task
	r.HandleFunc("POST /api/task/done", auth(h.doneTask))

	//delete task
	r.HandleFunc("DELETE /api/task", auth(h.deleteTask))

	//authorization
	r.HandleFunc("POST /api/signin", h.password)

	return r
}
