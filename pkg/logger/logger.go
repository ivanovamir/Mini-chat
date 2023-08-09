package logger

import (
	"log/slog"
	"os"
	"path"
)

type Logger struct {
	*slog.Logger
	outputPath,
	appVersion string
	mode int
}

func NewLogger(options ...Option) *Logger {
	l := new(Logger)
	for _, option := range options {
		option(l)
	}

	if err := os.MkdirAll(path.Dir(l.outputPath), 0755); err != nil {
		return nil
	}

	logFile, err := os.OpenFile(l.outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return nil
	}

	l.Logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
	})).With(slog.String("version", l.appVersion))

	return l
}
