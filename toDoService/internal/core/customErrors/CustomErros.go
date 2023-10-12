package customErrors

type UnauthorizedError struct {
}

func (s *UnauthorizedError) Error() string {
	return "Unauthorized!"
}
