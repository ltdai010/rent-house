package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/commentcontroller:CommentController"] = append(beego.GlobalControllerRouter["rent-house/controllers/commentcontroller:CommentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:commentID/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/commentcontroller:CommentController"] = append(beego.GlobalControllerRouter["rent-house/controllers/commentcontroller:CommentController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:commentID/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/commentcontroller:CommentController"] = append(beego.GlobalControllerRouter["rent-house/controllers/commentcontroller:CommentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:commentID/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
