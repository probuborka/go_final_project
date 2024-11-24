package config

import (
	"os"
	"path/filepath"

	entityconfig "github.com/probuborka/go_final_project/internal/entity/config"
)

type Config struct {
	HTTP entityconfig.HTTPConfig
	DB   entityconfig.DBConfig
	Auth entityconfig.Authentication
}

func New() (*Config, error) {
	// password
	password := os.Getenv("TODO_PASSWORD")

	//port
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = entityconfig.Port
	}

	//db
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = filepath.Join(entityconfig.DBDir, "/", entityconfig.DBName)
	}
	dbDriver := entityconfig.DBDriver
	dbCreate := entityconfig.DBCreate

	return &Config{
		HTTP: entityconfig.HTTPConfig{Port: port},
		DB: entityconfig.DBConfig{
			Driver: dbDriver,
			File:   dbFile,
			Create: dbCreate,
		},
		Auth: entityconfig.Authentication{
			Password: password,
		},
	}, nil
}
