package ports

import (
	"authService/internal/core/domain"
	"authService/internal/core/dtos"
)

type AuthService interface {
	AuthorizeUser(username string, password string) (string, error)
	ValidateToken(token string) (*dtos.UserDTO, error)
	RegisterUser(dtos.AddUserRequest) (token *string, err error)
}
type AuthRepository interface {
	FindUserById(id string) (*domain.User, error)
	FindUserByPassword(username string, password string) (*domain.User, error)
	AddUserToDb(user *domain.User) error
}

// i think i shouldnt make this as a service i will not change the implementation probablily. maybe using sessions over jwt tokens could be a change ?
type TokenService interface {
	GenerateToken(id string, username string, nameLastname string) (*string, error)
	ValidateToken(token string) (*dtos.UserDTO, error)
}
