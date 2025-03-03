package CustomErrors

import "fmt"

type CustomError struct {
	Message        string
	HttpStatusCode int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s (status: %d)", e.Message, e.HttpStatusCode)
}

func NewCustomError(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:        message,
		HttpStatusCode: statusCode,
	}
}

var (
	ErrUserNotFound       = NewCustomError("User couldn't be found!", 403)
	ErrInvalidToken       = NewCustomError("Invalid token!", 401)
	ErrUnauthorized       = NewCustomError("Unauthorized access!", 403)
	ErrInternalError      = NewCustomError("Internal server error!", 500)
	TokenIsNotValidError  = NewCustomError("Token is not valid!", 400)
	EmailOrUsernameExists = NewCustomError("this email or username already exists", 400)
)
