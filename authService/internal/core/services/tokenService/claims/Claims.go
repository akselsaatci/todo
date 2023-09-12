package claims

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId           string `json:"userId"`
	UserName         string `json:"userName"`
	UserNameLastname string `json:"userNameLastname"`
	jwt.StandardClaims
}
