package CustomErrors

type UserNotFoundError struct {
}

func (e *UserNotFoundError) Error() string {
	return "User couldn't found!"
}

type InvalidTokenError struct {
}

func (e *InvalidTokenError) Error() string {
	return "Invalid token!"
}
