package repository

import (
	"context"
	"database/sql"
	"strconv"

	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
	entitytask "github.com/probuborka/go_final_project/internal/entity/task"
)

type repoTask struct {
	db *sql.DB
}

func newRepoTask(db *sql.DB) repoTask {
	return repoTask{
		db: db,
	}
}

func (r repoTask) Create(ctx context.Context, task entitytask.Task) (int, error) {

	res, err := r.db.Exec(
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

func (r repoTask) Update(ctx context.Context, task entitytask.Task) error {

	res, err := r.db.Exec(
		`UPDATE scheduler 
		 SET date    = :date,
		     title   = :title,
			 comment = :comment,
			 repeat  = :repeat
		 WHERE id = :id`,
		sql.Named("id", task.ID),
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r repoTask) Get(ctx context.Context, search string, searchDate string) ([]entitytask.Task, error) {

	search = "%" + search + "%"

	rows, err := r.db.Query(
		`SELECT id, date, title, comment, repeat 
		 FROM scheduler 
		 WHERE title LIKE :search 
		    OR comment LIKE :search
			OR date = :searchDate
		 ORDER BY date 
		 LIMIT :limit`,
		sql.Named("search", search),
		sql.Named("searchDate", searchDate),
		sql.Named("limit", entityconfig.RowsLimit),
	)
	if err != nil {
		return nil, err
	}

	task := entitytask.Task{}
	tasks := make([]entitytask.Task, 0)

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r repoTask) GetById(ctx context.Context, idStr string) (entitytask.Task, error) {
	task := entitytask.Task{}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return task, err
	}

	row := r.db.QueryRow(
		`SELECT id, date, title, comment, repeat 
		 FROM scheduler 
		 WHERE id = :id`,
		sql.Named("id", id),
	)

	err = row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r repoTask) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(
		`DELETE FROM scheduler 
		 WHERE id = :id`,
		sql.Named("id", id),
	)
	if err != nil {
		return err
	}

	return nil
}
