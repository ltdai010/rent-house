package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
	"strings"
)

func filterRenter(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/renter/login") || strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/renter/sign-up") ||
		strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/renter/info") || isRenter(ctx)  {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func isRenter(ctx *context.Context) bool {
	token, err := jwt.ParseWithClaims(ctx.Input.Header("token"), &TokenClaims{}, KeyFunc)
	if err != nil {
		log.Println(err)
		return false
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		renter := &models.Renter{}
		err = renter.GetFromKey(claim.Username)
		if err != nil {
			log.Println(err)
			return false
		}
		if claim.IssuedAt < renter.PasswordChanged {
			return false
		}
		ctx.Request.Header.Set("rentername", claim.Username)
		return true
	}
	return false
}