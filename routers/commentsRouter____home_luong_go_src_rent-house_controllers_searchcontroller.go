package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/searchcontroller:SearchController"] = append(beego.GlobalControllerRouter["rent-house/controllers/searchcontroller:SearchController"],
        beego.ControllerComments{
            Method: "GetPageActivateSearchHouse",
            Router: "/page-search-results",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
