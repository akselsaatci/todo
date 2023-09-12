package Requests

type GetUserTodosRequest struct {
	Token string `json:"token" xml:"token" binding:"required"`
}
