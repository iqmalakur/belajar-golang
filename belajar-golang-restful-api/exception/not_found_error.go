package exception

type NotFoundError struct {
	message string
}

func (error NotFoundError) Error() string {
	return error.message
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{message: message}
}
