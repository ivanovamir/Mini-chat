package logger

import "mini_chat/internal/configs"

type Option func(*Logger)

func WithCfg(cfg *configs.LoggerConfig) Option {
	return func(l *Logger) {
		l.cfg = cfg
	}
}

func WithAppVersion(appVersion string) Option {
	return func(l *Logger) {
		l.appVersion = appVersion
	}
}
