package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"] = append(beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/most-view-this-month/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"] = append(beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"],
        beego.ControllerComments{
            Method: "GetHouseInLocation",
            Router: "/number-house-in-location/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"] = append(beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"],
        beego.ControllerComments{
            Method: "GetTimelineThisMonth",
            Router: "/timeline-this-month/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"] = append(beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"],
        beego.ControllerComments{
            Method: "GetViewByPrice",
            Router: "/view-by-price/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"] = append(beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"],
        beego.ControllerComments{
            Method: "GetViewInLocation",
            Router: "/view-in-location/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
