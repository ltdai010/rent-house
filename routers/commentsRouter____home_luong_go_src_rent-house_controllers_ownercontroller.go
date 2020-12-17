package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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
            Method: "GetOwner",
            Router: "/:ownerID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("ownerID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "GetAllHouse",
            Router: "/:ownerID/houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("ownerID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "GetPageHouse",
            Router: "/:ownerID/page-houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("ownerID", param.IsRequired, param.InPath),
				param.New("page", param.IsRequired),
				param.New("count", param.IsRequired),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "CreateHouse",
            Router: "/house/",
            AllowHTTPMethods: []string{"post"},
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
            Method: "GetAllNotice",
            Router: "/notification/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("page"),
				param.New("length"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"] = append(beego.GlobalControllerRouter["rent-house/controllers/ownercontroller:OwnerController"],
        beego.ControllerComments{
            Method: "ChangePass",
            Router: "/password",
            AllowHTTPMethods: []string{"put"},
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
