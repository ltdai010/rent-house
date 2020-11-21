package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: "/add-comment/:houseID",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "AddHouseToFavorite",
            Router: "/like/:houseID",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"] = append(beego.GlobalControllerRouter["rent-house/controllers/rentercontroller:RenterController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/sign-up/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
