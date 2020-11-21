package main

import (
	"rent-house/middlewares"
	"rent-house/models"
	_ "rent-house/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	middlewares.InitFilter()
	models.InitDataBase()
	beego.Run()
}
