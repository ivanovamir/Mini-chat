package app

import (
	"context"
	"fmt"
	"log"
	"mini_chat/internal/config"
	"mini_chat/internal/handler"
	"mini_chat/pkg/logger"
	"mini_chat/pkg/postgres"
)

func Run() {
	//TODO: implement config +
	//TODO: implement logger (slog) +
	//TODO: implement errors
	//TODO: implement handler
	//TODO: implement server +
	//TODO: implement db +

	cfg := config.NewConfig()

	if cfg == nil {
		//TODO: handle error
		log.Fatal(fmt.Errorf("failed to load config"))
		return
	}

	lg := logger.NewLogger(logger.WithCfg(&cfg.LoggerConfig), logger.WithAppVersion(cfg.AppVersion.AppVersion))

	if lg == nil {
		//TODO: handle error
		log.Fatal(fmt.Errorf("failed to set up logger"))
		return
	}

	lg.Info("info")

	db, err := postgres.NewPostgresDB(context.Background(), &cfg.PostgresDBConfig)

	if err != nil {
		lg.Error(err.Error())
		return
	}

	defer db.Close()

	srv := handler.NewHandler(&cfg.HandlerConfig, lg)

	if err := srv.InitRoutes(); err != nil {
		lg.Error(err.Error())
		return
	}
}
