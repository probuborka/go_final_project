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

func (h handler) Init() http.Handler {
	r := http.NewServeMux()

	//next date
	r.HandleFunc("GET /api/nextdate", h.getNextDate)

	//web
	r.Handle("/", http.FileServer(http.Dir(entity.WebDir)))

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

	//authorization
	r.HandleFunc("POST /api/signin", h.password)

	hand := logging(r)
	hand = auth(hand)

	//

	return hand
}
