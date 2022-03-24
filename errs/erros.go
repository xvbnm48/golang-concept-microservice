package errs

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{e.Code, e.Message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{404, message}
}

func NewUnexpedtedError(message string) *AppError {
	return &AppError{500, message}
}
