package middlewares

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func filterAdmin(ctx *context.Context) {
	if validAdmin(ctx.Input.Header("key")) {
		return
	}
	ctx.ResponseWriter.WriteHeader(403)
}

func validAdmin(key string) bool {
	if key == beego.AppConfig.String("admin::key") {
		return true
	}
	return false
}


