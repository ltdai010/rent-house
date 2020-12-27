package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/notificationcontroller:WebsocketController"] = append(beego.GlobalControllerRouter["rent-house/controllers/notificationcontroller:WebsocketController"],
        beego.ControllerComments{
            Method: "Join",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
