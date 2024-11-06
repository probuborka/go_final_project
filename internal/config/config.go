package config

import (
	"os"
	"path/filepath"

	"github.com/probuborka/go_final_project/internal/entity"
)

type Config struct {
	HTTP entity.HTTPConfig
	DB   entity.DBConfig
}

func New() (*Config, error) {
	//port
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = entity.Port
	}

	//db
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = filepath.Join(entity.DBDir, "/", entity.DBName)
	}
	dbDriver := entity.DBDriver
	dbCreate := entity.DBCreate

	return &Config{
		HTTP: entity.HTTPConfig{Port: port},
		DB: entity.DBConfig{
			Driver: dbDriver,
			File:   dbFile,
			Create: dbCreate,
		},
	}, nil
}
