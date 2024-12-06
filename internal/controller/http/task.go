package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
	entitytask "github.com/probuborka/go_final_project/internal/entity/task"
	"github.com/probuborka/go_final_project/pkg/logger"
)

type serviceTask interface {
	Create(ctx context.Context, task entitytask.Task) (int, error)
	Change(ctx context.Context, task entitytask.Task) error
	Get(ctx context.Context, search string) ([]entitytask.Task, error)
	GetById(ctx context.Context, id string) (entitytask.Task, error)
	Done(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
	NextDate(nowDate time.Time, date string, repeat string) (string, error)
}

func (h handler) createTask(w http.ResponseWriter, r *http.Request) {
	var task entitytask.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	id, err := h.task.Create(r.Context(), task)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	//
	response(w, entitytask.IdTask{ID: strconv.Itoa(id)}, http.StatusCreated)
}

func (h handler) getTasks(w http.ResponseWriter, r *http.Request) {

	search := r.FormValue("search")

	tasks, err := h.task.Get(r.Context(), search)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	response(w, &entitytask.Tasks{Tasks: tasks}, http.StatusOK)
}

func (h handler) getTask(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	task, err := h.task.GetById(r.Context(), id)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	response(w, task, http.StatusOK)
}

func (h handler) changeTask(w http.ResponseWriter, r *http.Request) {
	var task entitytask.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &task)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	err = h.task.Change(r.Context(), task)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	//
	response(w, struct{}{}, http.StatusCreated)
}

func (h handler) doneTask(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	err := h.task.Done(r.Context(), id)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	response(w, struct{}{}, http.StatusOK)
}

func (h handler) deleteTask(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	err := h.task.Delete(r.Context(), id)
	if err != nil {
		//
		response(w, entityerror.Error{Error: err.Error()}, http.StatusBadRequest)
		//
		logger.Error(err)
		return
	}

	response(w, struct{}{}, http.StatusOK)
}
