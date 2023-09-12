package Requests

import "github.com/google/uuid"

type UpdateTodoRequest struct {
	TodoId      uuid.UUID `json:"todoId" xml:"todoId" binding:"required"`
	Token       string    `json:"token" xml:"token" binding:"required"`
	Title       string    `json:"title" xml:"title" binding:"required"`
	Description string    `json:"description" xml:"description" binding:"required"`
	IsDone      bool      `json:"isDone" xml:"isDone" binding:"required"`
}
