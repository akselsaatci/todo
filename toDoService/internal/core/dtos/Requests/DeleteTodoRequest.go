package Requests

import "github.com/google/uuid"

type DeleteTodoRequest struct {
	Token  string    `json:"token" xml:"token" binding:"required"`
	TodoId uuid.UUID `json:"todoId" xml:"todoId" binding:"required"`
}
