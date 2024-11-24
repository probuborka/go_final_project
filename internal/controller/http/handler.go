package http

import (
	"net/http"

	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
)

var (
	cfg entityconfig.Authentication
)

type handler struct {
	task           serviceTask
	authentication serviceAuthentication
}

func New(task serviceTask, authentication serviceAuthentication, cfgAuth entityconfig.Authentication) *handler {
	cfg = cfgAuth
	return &handler{
		task:           task,
		authentication: authentication,
	}
}

func (h handler) Init() http.Handler {
	r := http.NewServeMux()

	//web
	r.Handle("/", http.FileServer(http.Dir(entityconfig.WebDir)))

	//next date
	r.HandleFunc("GET /api/nextdate", h.getNextDate)

	//create task
	r.HandleFunc("POST /api/task", h.createTask)

	//get tasks
	r.HandleFunc("GET /api/tasks", h.getTasks)

	//get task
	r.HandleFunc("GET /api/task", h.getTask)

	//change task
	r.HandleFunc("PUT /api/task", h.changeTask)

	//done task
	r.HandleFunc("POST /api/task/done", h.doneTask)

	//delete task
	r.HandleFunc("DELETE /api/task", h.deleteTask)

	//authentication
	r.HandleFunc("POST /api/signin", h.password)

	//
	stack := []middleware{
		logging,
		authentication,
	}

	hand := compileMiddleware(r, stack)

	return hand
}
