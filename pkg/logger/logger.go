package logger

import (
	"log/slog"
	"mini_chat/internal/configs"
	"os"
	"path"
)

type Logger struct {
	*slog.Logger
	cfg        *configs.LoggerConfig
	appVersion string
}

func NewLogger(options ...Option) *Logger {
	l := new(Logger)

	for _, option := range options {
		option(l)
	}

	if err := os.MkdirAll(path.Dir(l.cfg.Path), 0755); err != nil {
		return nil
	}

	logFile, err := os.OpenFile(l.cfg.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return nil
	}

	handlerOptions := &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: nil,
	}

	switch l.cfg.Level {
	case 1:
		handlerOptions.Level = slog.LevelError
	case 2:
		handlerOptions.Level = slog.LevelDebug
	}

	l.Logger = slog.New(slog.NewJSONHandler(logFile, handlerOptions)).With(slog.String("ver.", l.appVersion))

	return l
}
