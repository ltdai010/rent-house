package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
)

func filterComment(ctx *context.Context) {
	if ctx.Input.Method() == "GET" {
		return
	} else if ownComment(ctx.Input.Header("token"), ctx.Input.Param(":commentID")) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func ownComment(tokenString, commentID string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, keyFunc)
	if err != nil {
		log.Println(err)
		return false
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		comment := &models.Comment{}
		err = comment.GetFromKey(commentID)
		if comment.OwnerID == claim.Username {
			return true
		}
	}
	return false
}
