package addresscontroller

import (
	"github.com/astaxie/beego"
	"rent-house/restapi/response"
	"rent-house/services/addressservice"
)

type AddressController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Success 200 {object} models.Province
// @Failure 403 :commentID is empty
// @router /provinces [get]
func (u *AddressController) GetProvince() {
	prs, err := addressservice.GetAllProvince()
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: prs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetDistrict
// @Description get all district in province
// @Param	provinceID	path	string	true	"provinceID"
// @Success 200 {object} models.District
// @Failure 403 :commentID is empty
// @router /:provinceID/districts [get]
func (u *AddressController) GetDistrict() {
	provinceID := u.Ctx.Input.Param(":provinceID")
	prs, err := addressservice.GetAllDistrict(provinceID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: prs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetCommune
// @Description get all commune in district
// @Param	districtID	path	string	true	"districtID"
// @Success 200 {object} models.District
// @Failure 403 :commentID is empty
// @router /:districtID/communes [get]
func (u *AddressController) GetCommune() {
	districtID := u.Ctx.Input.Param(":districtID")
	prs, err := addressservice.GetAllCommune(districtID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: prs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}