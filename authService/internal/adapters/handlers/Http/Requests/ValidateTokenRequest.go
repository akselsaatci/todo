package Requests

type ValidateTokenRequest struct {
	Token string `json:"token" xml:"token" binding:"required"`
}
