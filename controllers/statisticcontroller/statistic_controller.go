package statisticcontroller

import (
	"github.com/astaxie/beego"
	"rent-house/restapi/response"
	"rent-house/services/statisticservice"
)

// Operations about Statistic
type StatisticController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	key		hear	string	true		"key"
// @Success 200 {object} models.Renter
// @Failure 403 :renterID is empty
// @router /most-view-this-month/ [get]
func (u *StatisticController) Get() {
	user, err := statisticservice.MostViewInMonth()
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: user,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}