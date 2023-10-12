package Requests

type CreateTodoRequest struct {
	Title       string `json:"title" xml:"title" binding:"required"`
	Description string `json:"description" xml:"description" binding:"required"`
}
