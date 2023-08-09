package main

import (
	"fmt"
	"mini_chat/internal/configs"
	"mini_chat/pkg/logger"
)

func main() {
	//TODO: implement config +
	//TODO: implement logger (slog) +
	//TODO: implement errors
	//TODO: implement handler
	//TODO: implement server
	//TODO: implement db

	cfg := configs.NewConfig()

	if cfg == nil {
		//TODO: handle error
	}

	lg := logger.NewLogger(logger.WithCfg(&cfg.LoggerConfig), logger.WithAppVersion(cfg.AppVersion.AppVersion))

	if lg == nil {
		//TODO: handle error
	}

	lg.Info("info")

	fmt.Println(cfg)
}
