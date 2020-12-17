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
            Method: "Delete",
            Router: "/:houseID/",
            AllowHTTPMethods: []string{"delete"},
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
            Method: "GetAllComment",
            Router: "/:houseID/comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "UpdateExpiredTime",
            Router: "/:houseID/expired-time",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("time", param.IsRequired),
			),
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
            Method: "GetByLike",
            Router: "/favorite-desc",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("page"),
				param.New("count"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"] = append(beego.GlobalControllerRouter["rent-house/controllers/housecontroller:HouseController"],
        beego.ControllerComments{
            Method: "UploadImage",
            Router: "/images",
            AllowHTTPMethods: []string{"post"},
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

}
