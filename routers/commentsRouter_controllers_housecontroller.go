package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetAllActivateHouse",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetAllSearchHouse",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:houseID/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:houseID/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:houseID/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetAllComment",
            Router: "/:houseID/comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetPageComment",
            Router: "/:houseID/page-comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetPageActivateHouse",
            Router: "/page",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "GetPageActivateSearchHouse",
            Router: "/search-results",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
