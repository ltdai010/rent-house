package middlewares

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/dgrijalva/jwt-go"
	"rent-house/consts"
	"time"
)

type TokenClaims struct {
	Username string
	jwt.StandardClaims
}

var(
	keyFunc = func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.SECRET_KEY), nil
	}
)

func InitFilter() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("/v1/rent-house/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Connection", "Authorization", "Sec-WebSocket-Extensions", "Sec-WebSocket-Key",
			"Sec-WebSocket-Version", "Access-Control-Allow-Origin", "content-type", "Content-Type", "sessionkey", "token", "Upgrade"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type", "Sec-WebSocket-Accept", "Connection", "Upgrade"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/v1/rent-house/renter/:renterID/*", beego.BeforeRouter, filterRenter)
	beego.InsertFilter("/v1/rent-house/owner/:ownerID/*", beego.BeforeRouter, filterOwner)
	beego.InsertFilter("/v1/rent-house/house/:houseID/*", beego.BeforeRouter, filterHouse)
	beego.InsertFilter("/v1/rent-house/comment/:commentID/*", beego.BeforeRouter, filterComment)
	beego.InsertFilter("/v1/rent-house/admin/*", beego.BeforeRouter, filterAdmin)
}

func CreateToken(username string) (string, error) {
	claims := TokenClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 360000,
			IssuedAt: time.Now().Unix(),
			Issuer:    "Test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(consts.SECRET_KEY))
}
