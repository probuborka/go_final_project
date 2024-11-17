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

func (t task) Get(ctx context.Context, search string) ([]entity.Task, error) {

	search = "%" + search + "%"

	rows, err := t.db.Query(
		`SELECT id, date, title, comment, repeat 
		 FROM scheduler 
		 WHERE title LIKE :search OR comment LIKE :search
		 ORDER BY date LIMIT :limit`,
		sql.Named("search", search),
		sql.Named("limit", 50),
	)
	if err != nil {
		return nil, err
	}

	tasks := make([]entity.Task, 0)
	v := entity.Task{}

	for rows.Next() {
		err := rows.Scan(&v.ID, &v.Date, &v.Title, &v.Comment, &v.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, v)
	}

	return tasks, nil
}
