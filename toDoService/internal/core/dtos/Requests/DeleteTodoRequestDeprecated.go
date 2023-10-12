package Requests

import "github.com/google/uuid"

type DeleteTodoRequest struct {
	TodoId uuid.UUID `json:"todoId" xml:"todoId" binding:"required"`
}
