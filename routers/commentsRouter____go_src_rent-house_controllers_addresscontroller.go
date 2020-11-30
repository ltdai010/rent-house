package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/addresscontroller:AddressController"] = append(beego.GlobalControllerRouter["rent-house/controllers/addresscontroller:AddressController"],
        beego.ControllerComments{
            Method: "GetCommune",
            Router: "/:districtID/communes",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/addresscontroller:AddressController"] = append(beego.GlobalControllerRouter["rent-house/controllers/addresscontroller:AddressController"],
        beego.ControllerComments{
            Method: "GetDistrict",
            Router: "/:provinceID/districts",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/addresscontroller:AddressController"] = append(beego.GlobalControllerRouter["rent-house/controllers/addresscontroller:AddressController"],
        beego.ControllerComments{
            Method: "GetProvince",
            Router: "/provinces",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
