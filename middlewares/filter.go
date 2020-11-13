package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"rent-house/consts"
	"rent-house/models"
	"time"
)

type TokenClaims struct {
	Login models.Login
	jwt.StandardClaims
}

var(
	keyFunc = func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.SECRET_KEY), nil
	}
)

func CreateToken(login models.Login) (string, error) {
	claims := TokenClaims{
		Login:       login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 360000,
			IssuedAt: time.Now().Unix(),
			Issuer:    "Test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(consts.SECRET_KEY))
}
