package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
	"strings"
)

func filterAdmin(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/v1/rent-house/admin/login") || ValidAdmin(ctx) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func ValidAdmin(ctx *context.Context) bool {
	token, err := jwt.ParseWithClaims(ctx.Input.Header("token"), &TokenClaims{}, KeyFunc)
	if err != nil {
		log.Println(err)
		return false
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		admin := &models.Admin{}
		err = admin.GetFromKey(claim.Username)
		if err != nil {
			return false
		}
		ctx.Request.Header.Set("admin", claim.Username)
		return true
	}
	return false
}

func GetAdminFromToken(tokenString string) string {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, KeyFunc)
	if err != nil {
		log.Println(err)
		return ""
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		admin := &models.Admin{}
		err = admin.GetFromKey(claim.Username)
		if err != nil {
			return ""
		}
		return claim.Username
	}
	return ""
}


