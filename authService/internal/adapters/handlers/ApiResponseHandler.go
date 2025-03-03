package handlers

import (
	"authService/internal/core/dtos"
	"net/http"
)

func BuildResponse[T any](responseStatus int, data T) dtos.ApiResponse[T] {
	return BuildResponse_(responseStatus, http.StatusText(responseStatus), data)
}

func BuildResponse_[T any](status int, message string, data T) dtos.ApiResponse[T] {
	return dtos.ApiResponse[T]{
		Message: message,
		Data:    data,
	}
}
