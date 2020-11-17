package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
)

func filterRenter(ctx *context.Context) {
	if isRenter(ctx.Input.Header("token"), ctx.Input.Param(":renterID")) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func isRenter(tokenString, renterID string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, keyFunc)
	if err != nil {
		log.Println(err)
		return false
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		renter := &models.Renter{}
		err = renter.GetFromKey(renterID)
		if renter.RenterName == claim.Username {
			return true
		}
	}
	return false
}