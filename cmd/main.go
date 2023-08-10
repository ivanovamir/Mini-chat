package main

import (
	"context"
	"fmt"
	"log"
	"mini_chat/internal/configs"
	"mini_chat/pkg/logger"
	"mini_chat/pkg/postgres"
)

func main() {
	//TODO: implement config +
	//TODO: implement logger (slog) +
	//TODO: implement errors
	//TODO: implement handler
	//TODO: implement server
	//TODO: implement db +

	cfg := configs.NewConfig()

	if cfg == nil {
		//TODO: handle error
	}

	lg := logger.NewLogger(logger.WithCfg(&cfg.LoggerConfig), logger.WithAppVersion(cfg.AppVersion.AppVersion))

	if lg == nil {
		//TODO: handle error
	}

	lg.Info("info")

	db, err := postgres.NewPostgresDB(context.Background(), &cfg.PostgresDBConfig)

	if err != nil {
		log.Printf("[ERROR]: %v", err.Error())
		return
	}

	defer db.Close()

	fmt.Println(cfg)
}
