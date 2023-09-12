package CustomErrors

type UserNotFoundError struct {
	Message string
	Code    string
}

func (e *UserNotFoundError) Error() string {
	return e.Code + " " + e.Message
}
