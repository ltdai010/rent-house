package housecontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rent-house/models"
	"rent-house/services/commentservices"
	"rent-house/services/houseservices"
	"rent-house/services/renterservices"
)

// Operations about house
type HouseController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all renters
// @Success 200 {object} models.Renter
// @router / [get]
func (u *HouseController) GetAll() {
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
// @Param	house-id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Renter
// @Failure 403 :house-id is empty
// @router /:house-id [get]
func (u *HouseController) Get() {
	id := u.Ctx.Input.Param(":house-id")
	if id != "" {
		user, err := houseservices.GetHouse(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title AddComment
// @Description create users
// @Param	body		body 	models.Comment	true		"body for user content"
// @Param	house-id	path	string			true		"the house id"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /:house-id/add-comment [post]
func (u *HouseController) AddComment() {
	var ob models.Comment
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	err = commentservices.AddComment(&ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	house-id		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Renter	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :house-id is not int
// @router /:house-id [put]
func (u *HouseController) Update() {
	id := u.Ctx.Input.Param(":house-id")
	if id != "" {
		var ob models.House
		err :=json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
		if err != nil {
			u.Ctx.WriteString(err.Error())
			return
		}
		err = houseservices.UpdateHouse(id, &ob)
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
// @Param	house-id		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 house-id is empty
// @router /:house-id [delete]
func (u *HouseController) Delete() {
	id := u.GetString(":house-id")
	err := houseservices.DeleteHouse(id)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
