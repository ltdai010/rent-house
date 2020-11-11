package admincontroler

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rent-house/models"
	"rent-house/services/ownerservices"
)

type AdminController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.Owner	true		"body for user content"
// @Success 200 {int} models.UserID
// @Failure 403 body is empty
// @router /create-owner [post]
func (u *AdminController) Post() {
	var user models.Owner
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	err = ownerservices.AddOwner(&user)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}
