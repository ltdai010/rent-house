package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/create-owner",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
