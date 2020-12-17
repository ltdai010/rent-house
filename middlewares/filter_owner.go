package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
	"strings"
)

func filterOwner(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/owner/login") || strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/owner/sign-up") || isOwner(ctx) || ctx.Input.Method() == "GET" {
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
		if ctx.Input.Method() == "PUT" {
			return true
		} else if !owner.Activate {
			return false
		}
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