package middlewares

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"rent-house/models"
)

func filterHouse(ctx *context.Context) {
	if ctx.Input.Method() == "GET" {
		return
	} else if ownHouse(ctx.Input.Header("token"), ctx.Input.Param(":houseID")) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func ownHouse(tokenString, houseID string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, KeyFunc)
	if err != nil {
		log.Println(err)
		return false
	}
	if claim, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		house := &models.House{}
		err = house.GetFromKey(houseID)
		if house.OwnerID == claim.Username {
			return true
		}
	}
	return false
}
