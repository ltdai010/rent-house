package test

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"rent-house/consts"
	"rent-house/models"
	"testing"
)

func TestAdmin(t *testing.T)  {
	models.InitDataBase()
	pass, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	_, err = models.Client.Collection(consts.ADMIN).Doc("admin").Set(context.Background(), map[string]string{
		"Username" : "admin",
		"Password" : string(pass),
	})
	if err != nil {
		t.Error(err)
	}
}
