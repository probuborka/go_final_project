package repository

import (
	"context"
	"database/sql"

	"github.com/probuborka/go_final_project/internal/entity"
)

type task struct {
	db *sql.DB
}

func newTask(db *sql.DB) task {
	return task{
		db: db,
	}
}

func (t task) Create(ctx context.Context, task entity.Task) (int, error) {

	res, err := t.db.Exec(
		`INSERT INTO scheduler (date, title, comment, repeat) 
			 VALUES (:date, :title, :comment, :repeat)`,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
