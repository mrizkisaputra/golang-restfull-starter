package exceptions

import "fmt"

type InternalServerError struct {
	message string
}

func (internalServerError InternalServerError) Error() string {
	return internalServerError.message
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{message: message}
}

func ErrorInternal(err error) {
	if err != nil {
		panic(NewInternalServerError(fmt.Sprintf("Internal Server Error: %s", err.Error())))
	}
}
