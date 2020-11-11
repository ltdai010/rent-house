package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:owner-id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:owner-id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:owner-id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "CreateHouse",
            Router: "/create-house",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/sign-up",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
