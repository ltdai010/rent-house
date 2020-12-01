// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"rent-house/controllers/addresscontroller"
	"rent-house/controllers/admincontroler"
	"rent-house/controllers/chatcontroller"
	"rent-house/controllers/commentcontroller"
	"rent-house/controllers/housecontroller"
	ownercontroller "rent-house/controllers/ownercontroller"
	"rent-house/controllers/rentercontroller"
	"rent-house/controllers/searchcontroller"
)

func init() {
	ns := beego.NewNamespace("/v1/rent-house",
		beego.NSNamespace("/owner",
			beego.NSInclude(
				&ownercontroller.OwnerController{},
				),
		),
		beego.NSNamespace("/house",
			beego.NSInclude(
				&housecontroller.HouseController{},
			),
		),
		beego.NSNamespace("/renter",
			beego.NSInclude(
				&rentercontroller.RenterController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&commentcontroller.CommentController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&admincontroler.AdminController{},
			),
		),
		beego.NSNamespace("/chat",
			beego.NSInclude(
				&chatcontroller.WebsocketController{},
			),
		),
		beego.NSNamespace("/address",
			beego.NSInclude(
				&addresscontroller.AddressController{},
			),
		),
		beego.NSNamespace("/search",
			beego.NSInclude(
				&searchcontroller.SearchController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
