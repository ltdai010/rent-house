package rentercontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rent-house/models"
	"rent-house/services/renterservices"
)

// Operations about Owner
type RenterController struct {
	beego.Controller
}

// @Title CreateRenter
// @Description create users
// @Param	body		body 	models.Renter	true		"body for user content"
// @Success 200 {int} models.UserID
// @Failure 403 body is empty
// @router /sign-up/ [post]
func (u *RenterController) Post() {
	var ob models.Renter
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	err = renterservices.AddRenter(&ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title CreateRenter
// @Description create users
// @Param	login		body 	models.Login	true		"body for user content"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /login/ [post]
func (u *RenterController) Login() {
	var ob models.Login
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	token, err := renterservices.LoginRenter(ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = token
	u.ServeJSON()
}

// @Title GetAll
// @Description get all renters
// @Success 200 {object} models.Renter
// @router / [get]
func (u *RenterController) GetAll() {
	users, err := renterservices.GetAllRenter()
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	token			header	string	true		"token"
// @Param	renterID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Renter
// @Failure 403 :renterID is empty
// @router /:renterID/ [get]
func (u *RenterController) Get() {
	id := u.Ctx.Input.Param(":renterID")
	if id != "" {
		user, err := renterservices.GetRenter(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	token		header		string	true		"The token"
// @Param	renterID	path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Renter	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :renterID is not int
// @router /:renterID/ [put]
func (u *RenterController) Put() {
	id := u.Ctx.Input.Param(":renterID")
	if id != "" {
		var ob models.Renter
		err :=json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
		if err != nil {
			u.Ctx.WriteString(err.Error())
			return
		}
		err = renterservices.UpdateRenter(id, &ob)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "success"
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	renterID		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 renterID is empty
// @router /:renterID/ [delete]
func (u *RenterController) Delete() {
	id := u.GetString(":renterID")
	err := renterservices.DeleteRenter(id)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
