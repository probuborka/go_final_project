package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/probuborka/go_final_project/internal/entity"
	"github.com/probuborka/go_final_project/pkg/logger"
)

type task interface {
	Create(ctx context.Context, task entity.Task) (int, error)
	Get(ctx context.Context, search string) ([]entity.Task, error)
}

func (h handler) createTask(w http.ResponseWriter, r *http.Request) {
	var jsonTask entity.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &jsonTask)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	id, err := h.task.Create(r.Context(), jsonTask)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	//
	response(w, entity.IdTask{ID: strconv.Itoa(id)}, http.StatusCreated)
}

func (h handler) getTasks(w http.ResponseWriter, r *http.Request) {

	search := r.FormValue("search")

	tasks, err := h.task.Get(r.Context(), search)
	if err != nil {
		//
		response(w, entity.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	response(w, &entity.Tasks{Tasks: tasks}, http.StatusOK)
}
