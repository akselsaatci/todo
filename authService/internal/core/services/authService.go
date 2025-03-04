package services

import (
	"authService/internal/core/domain"
	"authService/internal/core/dtos"
	CustomErrors "authService/internal/core/errors"
	"authService/internal/core/ports"
	"authService/internal/helper"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	tokenSerivce ports.TokenService
	repo         ports.AuthRepository
}

func NewAuthService(repo ports.AuthRepository, tokenService ports.TokenService) *AuthService {
	return &AuthService{repo: repo, tokenSerivce: tokenService}
}

func (a *AuthService) AuthorizeUser(username string, password string) (string, error) {
	user, err := a.repo.FindUserByPassword(username, password)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", CustomErrors.ErrUserNotFound
	}

	token, err := a.tokenSerivce.GenerateToken(user.ID.String(), user.UserName, user.NameLastName)
	if err != nil {
		return "", fmt.Errorf("Error happend while generating a token!")
	}

	return *token, nil
}

func (a *AuthService) ValidateToken(token string) (*dtos.UserDTO, error) {
	res, err := a.tokenSerivce.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res, nil
}
func (a *AuthService) RegisterUser(data dtos.AddUserRequest) (token *string, err error) {

	user := domain.User{
		ID:            uuid.New(),
		NameLastName:  data.NameLastname,
		UserName:      data.UserName,
		Email:         data.Email,
		Password:      helper.EncryptPassword(data.Password),
		LastLoginDate: time.Now(),
		IsVerified:    false,
	}

	err = a.repo.AddUserToDb(&user)
	if err != nil {
		return nil, err
	}
	token, err = a.tokenSerivce.GenerateToken(user.ID.String(), user.UserName, user.NameLastName)
	if err != nil {
		return nil, err
	}
	return token, nil
}
