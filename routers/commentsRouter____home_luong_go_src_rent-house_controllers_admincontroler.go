package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetAllComment",
            Router: "/:houseID/comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetPageComment",
            Router: "/:houseID/page-comments/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetReportInHouse",
            Router: "/:houseID/reports/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("page", param.IsRequired),
				param.New("length", param.IsRequired),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: "/:ownerID/messages/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
            Method: "DeniedExtendingHouse",
            Router: "/denied-extending-house/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "DeniedHouse",
            Router: "/denied-house/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetAllDeniedHouse",
            Router: "/denied-houses/",
            AllowHTTPMethods: []string{"get"},
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
            Method: "GetAllExtendHouse",
            Router: "/extend-houses/",
            AllowHTTPMethods: []string{"get"},
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
            Method: "DeactivateOwner",
            Router: "/inactive-owner/",
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
            Method: "GetMessagingOwner",
            Router: "/messages/owner",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "SendMessageToOwner",
            Router: "/messages/owner",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "DeleteOwner",
            Router: "/owner/",
            AllowHTTPMethods: []string{"delete"},
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
            Method: "GetPageDeniedHouse",
            Router: "/page-denied-houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetPageExtendHouse",
            Router: "/page-extend-houses/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetPageOwner",
            Router: "/page-owner",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("page", param.IsRequired),
				param.New("length", param.IsRequired),
			),
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
            Method: "DeleteReport",
            Router: "/report/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"] = append(beego.GlobalControllerRouter["rent-house/controllers/admincontroler:AdminController"],
        beego.ControllerComments{
            Method: "GetReport",
            Router: "/reports/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("page", param.IsRequired),
				param.New("length", param.IsRequired),
				param.New("status", param.IsRequired),
			),
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
