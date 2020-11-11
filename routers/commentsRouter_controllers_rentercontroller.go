package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:renter-id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:renter-id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:renter-id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/sign-up",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
