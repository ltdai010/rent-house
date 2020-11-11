// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"rent-house/controllers/admincontroler"
	"rent-house/controllers/commentcontroller"
	"rent-house/controllers/housecontroller"
	"rent-house/controllers/ownercontroller"
	"rent-house/controllers/rentercontroller"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/owner",
			beego.NSInclude(
				&OwnerController.OwnerController{},
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
	)
	beego.AddNamespace(ns)
}
