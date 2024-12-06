package repository

import "database/sql"

type repository struct {
	Task repoTask
}

func New(db *sql.DB) *repository {
	return &repository{
		Task: newRepoTask(db),
	}
}
