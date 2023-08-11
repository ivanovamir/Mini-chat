package logger

type Option func(*Logger)

func WithCfg(cfg *LoggerConfig) Option {
	return func(l *Logger) {
		l.cfg = cfg
	}
}

func WithAppVersion(appVersion string) Option {
	return func(l *Logger) {
		l.appVersion = appVersion
	}
}
