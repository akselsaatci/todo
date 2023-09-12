package tokenService

import (
	"authService/internal/core/dtos"
	"authService/internal/core/services/tokenService/claims"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenService struct {
	jwtKey []byte
}

func NewTokenService(_jwtKey []byte) *TokenService {
	return &TokenService{jwtKey: _jwtKey}
}

func (t *TokenService) GenerateToken(id string, username string, nameLastname string) (*string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	var jwtKey = t.jwtKey
	claims := &claims.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		UserId:           id,
		UserName:         username,
		UserNameLastname: nameLastname,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}
	return &tokenString, nil

}
func (t *TokenService) ValidateToken(token string) (*dtos.UserDTO, error) {
	user := dtos.UserDTO{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return t.jwtKey, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &claims.Claims{}, keyFunc)
	if err != nil {
		verr := err.(*jwt.ValidationError)

		return nil, errors.New(verr.Inner.Error())
	}

	payload, ok := jwtToken.Claims.(*claims.Claims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	user.ID = payload.UserId
	user.NameLastName = payload.UserNameLastname
	user.Username = payload.UserName

	return &user, nil

}
