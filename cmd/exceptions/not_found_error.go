package exceptions

type NotFoundError struct {
	message string
}

func (notFoundError NotFoundError) Error() string {
	return notFoundError.message
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{message: message}
}
