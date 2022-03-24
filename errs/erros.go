package errs

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{404, message}
}

func NewUnexpedtedError(message string) *AppError {
	return &AppError{500, message}
}
