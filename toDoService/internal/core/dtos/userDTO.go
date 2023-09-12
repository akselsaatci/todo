package dtos

type UserDTO struct {
	ID           string `json:"userId"`
	NameLastName string `json:"userName"`
	Username     string `json:"userNameLastname"`
	Email        string `json:"email"`
}
