package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/probuborka/go_final_project/internal/entity"
)

var (
	errValidateTask = errors.New("func validateTask()")
)

type dbTask interface {
	Create(ctx context.Context, task entity.Task) (int, error)
	Get(ctx context.Context, search string) ([]entity.Task, error)
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
		task.Date, err = entity.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			return 0, err
		}
	}

	id, err := t.db.Create(ctx, task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t task) Get(ctx context.Context, search string) ([]entity.Task, error) {

	tasks, err := t.db.Get(ctx, search)
	if err != nil {
		return nil, err
	}

	return tasks, nil
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
