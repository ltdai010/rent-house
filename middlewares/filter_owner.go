package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
	"strings"
)

func filterOwner(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/owner/login/") || strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/owner/sign-up/") || isOwner(ctx) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func isOwner(ctx *context.Context) bool {
	token, err := jwt.ParseWithClaims(ctx.Input.Header("token"), &TokenClaims{}, keyFunc)
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
		ctx.Request.Header.Set("ownername", claim.Username)
		return true
	}
	return false
}
