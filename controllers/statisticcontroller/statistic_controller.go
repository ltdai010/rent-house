package statisticcontroller

import (
	"github.com/astaxie/beego"
	"rent-house/restapi/response"
	"rent-house/services/statisticservices"
)

// Operations about Statistic
type StatisticController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	key		header	string	true		"key"
// @Param	length	query	int true		"page number"
// @Success 200 {object} response.House
// @Failure 403 :houseID is empty
// @router /most-view-this-month/ [get]
func (u *StatisticController) Get() {
	length, _ := u.GetInt("length")
	user, err := statisticservices.MostViewInMonth(length)
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


// @Title GetTimelineThisMonth
// @Description get view in hour
// @Param	key		header	string	true		"key"
// @Success 200 {map} map[string]int64{}
// @Failure 403 : is empty
// @router /timeline-this-month/ [get]
func (u *StatisticController) GetTimelineThisMonth() {
	stat, err := statisticservices.TimelineViewThisMonth()
	if err != nil {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: stat,
			Err:  response.NewErr(response.NotExisted),
		}
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: stat,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}