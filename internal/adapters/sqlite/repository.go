package repository

import "database/sql"

type repository struct {
	Task task
}

func New(db *sql.DB) *repository {
	return &repository{
		Task: newTask(db),
	}
}
