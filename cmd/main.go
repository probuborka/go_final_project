package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	repository "github.com/probuborka/go_final_project/internal/adapters/sqlite"
	"github.com/probuborka/go_final_project/internal/config"
	handler "github.com/probuborka/go_final_project/internal/controller/http/v1"
	"github.com/probuborka/go_final_project/internal/service"
	"github.com/probuborka/go_final_project/pkg/logger"
	"github.com/probuborka/go_final_project/pkg/route"
	"github.com/probuborka/go_final_project/pkg/sqlite"
)

func main() {
	//config
	cfg, err := config.New()
	if err != nil {
		logger.Error(err)
		return
	}

	//db
	db, err := sqlite.New(cfg.DB.Driver, cfg.DB.File, cfg.DB.Create)
	if err != nil {
		logger.Error(err)
		return
	}
	defer db.Close()

	//repo
	repo := repository.New(db)

	//service
	service := service.New(repo.Task)

	//handlers
	handlers := handler.New(service.Task)

	//http server
	server := route.New(cfg.HTTP.Port, handlers.Init())

	//start server
	logger.Info("server start")
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	//stop server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Info("server stop")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
