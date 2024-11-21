package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/probuborka/go_final_project/internal/entity"
	"github.com/probuborka/go_final_project/internal/service/nextdate"
)

type repository interface {
	Create(ctx context.Context, task entity.Task) (int, error)
	Change(ctx context.Context, task entity.Task) error
	Get(ctx context.Context, search string, searchDate string) ([]entity.Task, error)
	GetById(ctx context.Context, id string) (entity.Task, error)
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

func (s service) Create(ctx context.Context, task entity.Task) (int, error) {

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

		if nowDate.Format(entity.Format1) > task.Date {
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

func (s service) Change(ctx context.Context, task entity.Task) error {
	//check
	if task.ID == "" {
		return entity.ErrNoID
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

		if nowDate.Format(entity.Format1) > task.Date {
			task.Date = date.Next()
		}
	}

	//db change task
	err = s.repo.Change(ctx, task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.ErrTaskNotFound
		}
		return err
	}

	return nil
}

func (s service) Get(ctx context.Context, search string) ([]entity.Task, error) {
	var searchDate string

	date, err := time.Parse(entity.Format2, search)
	if err == nil {
		searchDate = date.Format(entity.Format1)
	}

	//get change tasks
	tasks, err := s.repo.Get(ctx, search, searchDate)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s service) GetById(ctx context.Context, id string) (entity.Task, error) {
	//check
	if id == "" {
		return entity.Task{}, entity.ErrNoID
	}

	//get task by id
	task, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Task{}, entity.ErrTaskNotFound
		}
		return entity.Task{}, err
	}

	return task, nil
}

func (s service) Done(ctx context.Context, id string) error {
	// check
	if id == "" {
		return entity.ErrNoID
	}

	//done task
	task, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.ErrTaskNotFound
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
				return entity.ErrTaskNotFound
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
		return entity.ErrNoID
	}

	//get task by id
	_, err := s.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.ErrTaskNotFound
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

func validateTask(task *entity.Task) error {

	if strings.TrimSpace(task.Title) == "" {
		return fmt.Errorf("%w: Title", entity.ErrIsEmpty)
	}

	if task.Date == "" {
		task.Date = time.Now().Format(entity.Format1)
	}

	_, err := time.Parse(entity.Format1, task.Date)
	if err != nil {
		return fmt.Errorf("%w: Date", entity.ErrFormatError)
	}

	nowTime := time.Now().Format(entity.Format1)
	if task.Date < nowTime {
		task.Date = nowTime
	}

	return nil
}
