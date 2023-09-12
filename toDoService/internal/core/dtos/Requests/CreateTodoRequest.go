package Requests

type CreateTodoRequest struct {
	Token       string `json:"token" xml:"token" binding:"required"`
	Title       string `json:"title" xml:"title" binding:"required"`
	Description string `json:"description" xml:"description" binding:"required"`
}
