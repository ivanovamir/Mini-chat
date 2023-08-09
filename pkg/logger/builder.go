package logger

type Option func(*Logger)

func WithOutputPath(outputPath string) Option {
	return func(l *Logger) {
		l.outputPath = outputPath
	}
}

func WithAppVersion(appVersion string) Option {
	return func(l *Logger) {
		l.appVersion = appVersion
	}
}

func WithLoggerMode(mode int) Option {
	return func(l *Logger) {
		l.mode = mode
	}
}
