package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "CreateHouse",
            Router: "/create-house/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "GetAllHouse",
            Router: "/houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "GetPageHouse",
            Router: "/page-houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "CreateOwner",
            Router: "/sign-up/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
