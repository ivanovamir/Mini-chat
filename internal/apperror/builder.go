package apperror

type Option func(*AppError)

func WithError(err error) Option {
	return func(err *AppError) {
		err.Err = err
	}
}

func WithMessage(message string) Option {
	return func(err *AppError) {
		err.Message = message
	}
}

func WithDeveloperMessage(developerMessage string) Option {
	return func(err *AppError) {
		err.DeveloperMessage = developerMessage
	}
}

func WithCode(code string) Option {
	return func(err *AppError) {
		err.Code = code
	}
}
