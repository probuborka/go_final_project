package sqlite

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func New(dbDriver, dbFile, dbCreate string) (*sql.DB, error) {
	var install bool
	_, err := os.Stat(dbFile)
	if err != nil {
		install = true
	}

	db, err := sql.Open(dbDriver, dbFile)
	if err != nil {
		return nil, err
	}

	if install {
		_, err = db.Exec(dbCreate)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
