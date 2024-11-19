package service

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

var (
	errValidateTask = errors.New("func validateTask()")
)

type dbTask interface {
	Create(ctx context.Context, task entity.Task) (int, error)
	Change(ctx context.Context, task entity.Task) error
	Get(ctx context.Context, search string) ([]entity.Task, error)
	GetById(ctx context.Context, id string) (entity.Task, error)
	Delete(ctx context.Context, id string) error
}

type task struct {
	db dbTask
}

func newTask(db dbTask) task {
	return task{
		db: db,
	}
}

func (t task) Create(ctx context.Context, task entity.Task) (int, error) {

	err := validateTask(&task)
	if err != nil {
		return 0, err
	}

	if strings.TrimSpace(task.Repeat) != "" {
		nowDate := time.Now()

		date, err := nextdate.New(nowDate, task.Date, task.Repeat)
		if err != nil {
			return 0, err
		}

		if nowDate.Format(entity.Format) > task.Date {
			task.Date = date.Next()
		}
	}

	id, err := t.db.Create(ctx, task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t task) Change(ctx context.Context, task entity.Task) error {
	if task.ID == "" {
		return entity.ErrNoID
	}

	err := validateTask(&task)
	if err != nil {
		return err
	}

	if strings.TrimSpace(task.Repeat) != "" {
		nowDate := time.Now()

		date, err := nextdate.New(nowDate, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		if nowDate.Format(entity.Format) > task.Date {
			task.Date = date.Next()
		}
	}

	err = t.db.Change(ctx, task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.ErrTaskNotFound
		}
		return err
	}

	return nil
}

func (t task) Get(ctx context.Context, search string) ([]entity.Task, error) {

	tasks, err := t.db.Get(ctx, search)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t task) GetById(ctx context.Context, id string) (entity.Task, error) {
	//check
	if id == "" {
		return entity.Task{}, entity.ErrNoID
	}

	//get task
	task, err := t.db.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Task{}, entity.ErrTaskNotFound
		}
		return entity.Task{}, err
	}

	return task, nil
}

func (t task) Done(ctx context.Context, id string) error {
	// check
	if id == "" {
		return entity.ErrNoID
	}

	//get task
	task, err := t.db.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.ErrTaskNotFound
		}
		return err
	}

	//
	if strings.TrimSpace(task.Repeat) != "" {
		nowDate := time.Now()
		date, err := nextdate.New(nowDate, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		task.Date = date.Next()

		err = t.db.Change(ctx, task)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return entity.ErrTaskNotFound
			}
			return err
		}
	} else {
		err = t.db.Delete(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t task) Delete(ctx context.Context, id string) error {
	// check
	if id == "" {
		return entity.ErrNoID
	}

	//get task
	_, err := t.db.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.ErrTaskNotFound
		}
		return err
	}

	//delete task
	err = t.db.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func validateTask(task *entity.Task) error {

	if strings.TrimSpace(task.Title) == "" {
		return fmt.Errorf("%w: Title is empty", errValidateTask)
	}

	if task.Date == "" {
		task.Date = time.Now().Format(entity.Format)
	}

	_, err := time.Parse(entity.Format, task.Date)
	if err != nil {
		return fmt.Errorf("%w: Date format error", errValidateTask)
	}

	nowTime := time.Now().Format(entity.Format)
	if task.Date < nowTime {
		task.Date = nowTime
	}

	return nil
}
