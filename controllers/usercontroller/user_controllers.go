package usercontroller

import (
	"encoding/json"
	"rent-house/restapi/request"
	"rent-house/services/ownerservices"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	request.PostUser	true		"body for user content"
// @Success 200 {int} models.UserID
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user request.PostUser
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := ownerservices.AddUser(&user)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users, leng, err := ownerservices.GetAllUser()
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = map[string]interface{}{
		"list_user" : users,
		"length"	: leng,
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		user, err := ownerservices.GetUser(uid)
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
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	request.PutUser	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		var user request.PutUser
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		err := ownerservices.UpdateUser(&user)
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
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	err := ownerservices.DeleteUser(uid)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}