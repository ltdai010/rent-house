package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "ActivateComment",
            Router: "/active-comment/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "ActivateHouse",
            Router: "/active-house/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "ActivateOwner",
            Router: "/active-owner/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "ExtendHouse",
            Router: "/extend-house/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "CreateHouse",
            Router: "/house/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetAllOwner",
            Router: "/owners/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetPageWaitComment",
            Router: "/page-wait-comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetPageWaitHouse",
            Router: "/page-wait-houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetPageWaitOwner",
            Router: "/page-wait-owners/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetAllRenter",
            Router: "/renters/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetAllWaitComment",
            Router: "/wait-comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetAllWaitHouse",
            Router: "/wait-houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetAllWaitOwner",
            Router: "/wait-owners/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
