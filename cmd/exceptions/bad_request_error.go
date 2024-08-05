package exceptions

type BadRequestError struct {
	message string
}

func (badRequestError BadRequestError) Error() string {
	return badRequestError.message
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{message: message}
}
