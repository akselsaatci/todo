package dtos

type AddUserRequest struct {
	UserName     string `json:"userName" binding:"required"`
	Password     string `json:"password" binding:"required"`
	NameLastname string `json:"nameLastname" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
}
