package OwnerController

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
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
// @Param	files		formData	[]file			true		"house image"
// @Success 200 {int} models.House
// @Failure 403 body is empty
// @router /create-house/ [post]
func (u *OwnerController) CreateHouse() {
	var ob request.HousePost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	file, err := u.GetFiles("files")
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	if len(file) < 3 {
		u.Data["json"] = response.NewErr(response.UnSuccess)
		u.ServeJSON()
		return
	}
	ownerID := u.Ctx.Input.Header("username")
	s, err := houseservices.AddHouse(ownerID, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	err = houseservices.UploadFile(s, file)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title CreateOwner
// @Description create users
// @Param	body		body 	request.OwnerPost	true		"body for user content"
// @Success 200 {int} models.UserID
// @Failure 403 body is empty
// @router /sign-up/ [post]
func (u *OwnerController) CreateOwner() {
	var user request.OwnerPost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	err = ownerservices.AddOwner(&user)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}


// @Title Get
// @Description get user by uid
// @Param	token			header	string	true		"The token string"
// @Success 200 {object} models.User
// @Failure 403 :ownerID is empty
// @router / [get]
func (u *OwnerController) Get() {
	uid := u.Ctx.Input.Header("username")
	user, err := ownerservices.GetOwner(uid)
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

// @Title Update
// @Description update the user
// @Param	token			header	string	true		"The token string"
// @Param	body		body 	models.Owner	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :ownerID is not int
// @router / [put]
func (u *OwnerController) Put() {
	uid := u.Ctx.Input.Header("ownername")
	var user request.OwnerPut
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	err = ownerservices.UpdateOwner(uid, &user)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	token			header	string	true		"The token string"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router / [delete]
func (u *OwnerController) Delete() {
	uid := u.Ctx.Input.Header("ownername")
	err := ownerservices.DeleteOwner(uid)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title GetAllHouse
// @Description get all renters
// @Param	token			header	string	true		"The token string"
// @Success 200 {object} models.House
// @router /houses/ [get]
func (u *OwnerController) GetAllHouse() {
	id := u.Ctx.Input.Header("ownername")
	houses, err := houseservices.GetAllHouseOfOwner(id)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: houses,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageHouse
// @Description get page house
// @Param	token		header	string	true		"The token string"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.House
// @router /page-houses/ [get]
func (u *OwnerController) GetPageHouse() {
	id := u.Ctx.Input.Header("ownername")
	page, err := u.GetInt("page")
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	count, err := u.GetInt("count")
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	houses, err := houseservices.GetPageHouseOfOwner(id, page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: houses,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title Login
// @Description login
// @Param	login		body 	models.Login	true		"body for user content"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /login/ [post]
func (u *OwnerController) Login() {
	var ob models.Login
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	token, err := ownerservices.LoginOwner(ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: token,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}