package dtos

type UserDTO struct {
	ID           string `json:"userId" binding:"required"`
	NameLastName string `json:"nameLastname" binding:"required"`
	Username     string `json:"userName" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
}
