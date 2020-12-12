package adminservice

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"rent-house/middlewares"
	"rent-house/models"
	"rent-house/restapi/response"
)

func LoginAdmin(login models.Login) (string, error) {
	admin := &models.Admin{}
	err := admin.GetFromKey(login.Username)
	if err != nil {
		return "", err
	}
	log.Println(admin)
	if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(login.Password)) == nil {
		return middlewares.CreateToken(login.Username)
	}
	return "", response.NotAdmin
}
