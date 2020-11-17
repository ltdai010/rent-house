package OwnerController

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"rent-house/models"
	"rent-house/services/houseservices"
	"rent-house/services/ownerservices"
)

// Operations about Owner
type OwnerController struct {
	beego.Controller
}

// @Title CreateHouse
// @Description create users
// @Param	token		header	    string			true		"The token string"
// @Param	body		body 		models.House	true		"body for user content"
// @Param	ownerID		path		string			true		"owner id"
// @Param	files		formData	[]file			true		"house image"
// @Success 200 {int} models.House
// @Failure 403 body is empty
// @router /:ownerID/create-house/ [post]
func (u *OwnerController) CreateHouse() {
	var ob models.House
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	file, err := u.GetFiles("files")
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	if len(file) < 3 {
		u.Ctx.WriteString(errors.New("not enough image").Error())
		return
	}
	ownerID := u.Ctx.Input.Param(":ownerID")
	s, err := houseservices.AddHouse(ownerID, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	err = houseservices.UploadFile(s, file)
	u.Data["json"] = map[string]string{"ID" : s}
	u.ServeJSON()
}

// @Title CreateOwner
// @Description create users
// @Param	body		body 	models.Owner	true		"body for user content"
// @Success 200 {int} models.UserID
// @Failure 403 body is empty
// @router /sign-up/ [post]
func (u *OwnerController) CreateOwner() {
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


// @Title GetAll
// @Description get all owners
// @Success 200 {object} models.User
// @router / [get]
func (u *OwnerController) GetAll() {
	users, err := ownerservices.GetAllOwner()
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = map[string]interface{}{
		"list_user" : users,
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	token			header	string	true		"The token string"
// @Param	ownerID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :ownerID is empty
// @router /:ownerID/ [get]
func (u *OwnerController) Get() {
	uid := u.Ctx.Input.Param(":ownerID")
	if uid != "" {
		user, err := ownerservices.GetOwner(uid)
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
// @Param	token			header	string	true		"The token string"
// @Param	ownerID		path 	string	true		"The ownerID you want to update"
// @Param	body		body 	models.Owner	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :ownerID is not int
// @router /:ownerID/ [put]
func (u *OwnerController) Put() {
	uid := u.Ctx.Input.Param(":ownerID")
	if uid != "" {
		var user models.Owner
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		err := ownerservices.UpdateUser(uid, &user)
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
// @Param	token			header	string	true		"The token string"
// @Param	ownerID		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:ownerID/ [delete]
func (u *OwnerController) Delete() {
	uid := u.GetString(":ownerID")
	err := ownerservices.DeleteOwner(uid)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title GetAllHouse
// @Description get all renters
// @Param	ownerID	path	string	true	"the house-id
// @Success 200 {object} models.Renter
// @router /:ownerID/houses/ [get]
func (u *OwnerController) GetAllHouse() {
	id := u.Ctx.Input.Param(":ownerID")
	houses, err := houseservices.GetAllHouseOfOwner(id)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = houses
	u.ServeJSON()
}

// @Title GetPageHouse
// @Description get page house
// @Param	ownerID	path	string	true	"the house-id"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.House
// @router /:ownerID/page-houses/ [get]
func (u *OwnerController) GetPageHouse() {
	id := u.Ctx.Input.Param(":ownerID")
	page, err := u.GetInt("page")
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	count, err := u.GetInt("count")
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	users, err := houseservices.GetPageHouseOfOwner(id, page, count)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title CreateRenter
// @Description create users
// @Param	login		body 	models.Login	true		"body for user content"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /login/ [post]
func (u *OwnerController) Login() {
	var ob models.Login
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	token, err := ownerservices.LoginOwner(ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = token
	u.ServeJSON()
}