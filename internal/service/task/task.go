package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
	entityerror "github.com/probuborka/go_final_project/internal/entity/error"
	entitytask "github.com/probuborka/go_final_project/internal/entity/task"
	"github.com/probuborka/go_final_project/internal/service/nextdate"
)

type repository interface {
	Create(ctx context.Context, task entitytask.Task) (int, error)
	Change(ctx context.Context, task entitytask.Task) error
	Get(ctx context.Context, search string, searchDate string) ([]entitytask.Task, error)
	GetById(ctx context.Context, id string) (entitytask.Task, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo repository
}

func New(repo repository) service {
	return service{
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, task entitytask.Task) (int, error) {

	//check
	err := validateTask(&task)
	if err != nil {
		return 0, err
	}

	//repeat
	if strings.TrimSpace(task.Repeat) != "" {
		nowDate := time.Now()

		date, err := nextdate.New(nowDate, task.Date, task.Repeat)
		if err != nil {
			return 0, err
		}

		if nowDate.Format(entityconfig.Format1) > task.Date {
			task.Date = date.Next()
		}
	}

	//db create task
	id, err := s.repo.Create(ctx, task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s service) Change(ctx context.Context, task entitytask.Task) error {
	//check
	if task.ID == "" {
		return entityerror.ErrNoID
	}

	err := validateTask(&task)
	if err != nil {
		return err
	}

	//repeat
	if strings.TrimSpace(task.Repeat) != "" {
		nowDate := time.Now()

		date, err := nextdate.New(nowDate, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		if nowDate.Format(entityconfig.Format1) > task.Date {
			task.Date = date.Next()
		}
	}

	//db change task
	err = s.repo.Change(ctx, task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entityerror.ErrTaskNotFound
		}
		return err
	}

	return nil
}

func (s service) Get(ctx context.Context, search string) ([]entitytask.Task, error) {
	var searchDate string

	date, err := time.Parse(entityconfig.Format2, search)
	if err == nil {
		searchDate = date.Format(entityconfig.Format1)
	}

	//get change tasks
	tasks, err := s.repo.Get(ctx, search, searchDate)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s service) GetById(ctx context.Context, id string) (entitytask.Task, error) {
	//check
	if id == "" {
		return entitytask.Task{}, entityerror.ErrNoID
	}

	//get task by id
	task, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entitytask.Task{}, entityerror.ErrTaskNotFound
		}
		return entitytask.Task{}, err
	}

	return task, nil
}

func (s service) Done(ctx context.Context, id string) error {
	// check
	if id == "" {
		return entityerror.ErrNoID
	}

	//done task
	task, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entityerror.ErrTaskNotFound
		}
		return err
	}

	//repeat
	if strings.TrimSpace(task.Repeat) != "" {
		nowDate := time.Now()
		date, err := nextdate.New(nowDate, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		task.Date = date.Next()

		err = s.repo.Change(ctx, task)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return entityerror.ErrTaskNotFound
			}
			return err
		}
	} else {
		//delete task by id
		err = s.repo.Delete(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s service) Delete(ctx context.Context, id string) error {
	// check
	if id == "" {
		return entityerror.ErrNoID
	}

	//get task by id
	_, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entityerror.ErrTaskNotFound
		}
		return err
	}

	//delete task by id
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (t service) NextDate(nowDate time.Time, dateStr string, repeat string) (string, error) {

	date, err := nextdate.New(nowDate, dateStr, repeat)
	if err != nil {
		return "", err
	}

	return date.Next(), nil
}

func validateTask(task *entitytask.Task) error {

	if strings.TrimSpace(task.Title) == "" {
		return fmt.Errorf("%w: Title", entityerror.ErrIsEmpty)
	}

	if task.Date == "" {
		task.Date = time.Now().Format(entityconfig.Format1)
	}

	_, err := time.Parse(entityconfig.Format1, task.Date)
	if err != nil {
		return fmt.Errorf("%w: Date", entityerror.ErrFormatError)
	}

	nowTime := time.Now().Format(entityconfig.Format1)
	if task.Date < nowTime {
		task.Date = nowTime
	}

	return nil
}
