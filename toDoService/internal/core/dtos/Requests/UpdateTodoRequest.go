package Requests

type UpdateTodoRequest struct {
	TodoId      string `json:"todoId" xml:"todoId" binding:"required"`
	Title       string `json:"title" xml:"title" binding:"required"`
	Description string `json:"description" xml:"description" binding:"required"`
	IsDone      bool   `json:"isDone" xml:"isDone" binding:"required"`
}
