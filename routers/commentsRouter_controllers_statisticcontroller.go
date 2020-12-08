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
            Method: "GetTimelineThisMonth",
            Router: "/timeline-this-month/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"] = append(beego.GlobalControllerRouter["rent-house/controllers/statisticcontroller:StatisticController"],
        beego.ControllerComments{
            Method: "GetViewInHour",
            Router: "/view-in-hour-this-month/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
