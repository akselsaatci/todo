package dtos

type AddUserRequest struct {
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	NameLastname string `json:"nameLastname"`
	Email        string `json:"email"`
}
