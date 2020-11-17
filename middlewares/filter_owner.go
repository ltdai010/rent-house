package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
)

func filterOwner(ctx *context.Context) {
	if isOwner(ctx.Input.Header("token"), ctx.Input.Param(":ownerID")) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func isOwner(tokenString, ownerID string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, keyFunc)
	if err != nil {
		log.Println(err)
		return false
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		owner := &models.Owner{}
		err = owner.GetFromKey(ownerID)
		if owner.OwnerName == claim.Username {
			return true
		}
	}
	return false
}
