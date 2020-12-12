package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
	"strings"
)

func filterOwner(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/owner/login/") || strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/owner/sign-up/") || ctx.Input.Method() == "GET" || isOwner(ctx) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func isOwner(ctx *context.Context) bool {
	token, err := jwt.ParseWithClaims(ctx.Input.Header("token"), &TokenClaims{}, KeyFunc)
	if err != nil {
		log.Println(err)
		return false
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		owner := &models.Owner{}
		err = owner.GetFromKey(claim.Username)
		if err != nil {
			return false
		}
		if claim.IssuedAt < owner.PasswordChanged {
			return false
		}
		ctx.Request.Header.Set("ownername", claim.Username)
		return true
	}
	return false
}

func GetOwnernameFromToken(tokenString string) string {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, KeyFunc)
	if err != nil {
		return ""
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		owner := &models.Owner{}
		err = owner.GetFromKey(claim.Username)
		if err != nil {
			return ""
		}
		return claim.Username
	}
	return ""
}