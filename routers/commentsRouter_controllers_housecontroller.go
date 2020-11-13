package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:house-id/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:house-id/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:house-id/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: "/:house-id/add-comment",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetAllComment",
            Router: "/:house-id/comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetPageComment",
            Router: "/:house-id/page-comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
