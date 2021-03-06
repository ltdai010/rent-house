package rentercontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"rent-house/services/commentservices"
	"rent-house/services/houseservices"
	"rent-house/services/renterservices"
	"rent-house/services/reportservices"
)

// Operations about Renter
type RenterController struct {
	beego.Controller
}

// @Title CreateRenter
// @Description create users
// @Param	body		body 	request.RenterPost	true		"body for user content"
// @Success 200 {int} models.UserID
// @Failure 403 body is empty
// @router /sign-up/ [post]
func (u *RenterController) Post() {
	var ob request.RenterPost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	err = renterservices.AddRenter(&ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.Existed)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title Login
// @Description Login
// @Param	login		body 	models.Login	true		"body for user content"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /login/ [post]
func (u *RenterController) Login() {
	var ob models.Login
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	token, err := renterservices.LoginRenter(ob)
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.ErrLogin)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: token,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title AddComment
// @Description create comment
// @Param	token			header			string				true		"The token string"
// @Param	houseID			path			string				true		"the house id"
// @Param	body			body 			request.CommentPost	true		"body for user content"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /comment/:houseID [post]
func (u *RenterController) AddComment() {
	var ob request.CommentPost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	houseID := u.Ctx.Input.Param(":houseID")
	renterID := u.Ctx.Input.Header("rentername")
	err = commentservices.AddComment(houseID, renterID, &ob)
	if err != nil {

		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title AddReport
// @Description create comment
// @Param	token			header			string				true		"The token string"
// @Param	houseID			path			string				true		"the house id"
// @Param	body			body 			request.ReportPost	true		"body for user content"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /report/:houseID [post]
func (u *RenterController) AddReport() {
	var ob request.ReportPost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	houseID := u.Ctx.Input.Param(":houseID")
	renterID := u.Ctx.Input.Header("rentername")
	err = reportservices.AddReport(houseID, renterID, ob)
	if err != nil {

		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title AddOrRemoveHouseFromFavorite
// @Description create comment
// @Param	token		header	string	true		"The token string"
// @Param	houseID		path	string			true		"the house id"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /like/:houseID [put]
func (u *RenterController) AddOrRemoveHouseFromFavorite() {
	houseID := u.Ctx.Input.Param(":houseID")
	renterID := u.Ctx.Input.Header("rentername")
	err := houseservices.AddOrRemoveFromFavourite(renterID, houseID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}


// @Title Get
// @Description get user by uid
// @Param	token			header	string	true		"token"
// @Success 200 {object} models.Renter
// @Failure 403 :renterID is empty
// @router / [get]
func (u *RenterController) Get() {
	id := u.Ctx.Input.Header("rentername")
	user, err := renterservices.GetRenter(id)
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

// @Title GetInfo
// @Description get user by uid
// @Param	renterID		path	string	true		"renter id"
// @Success 200 {object} response.RenterInfo
// @Failure 403 :renterID is empty
// @router /info/:renterID [get]
func (u *RenterController) GetInfo() {
	id := u.Ctx.Input.Param(":renterID")
	user, err := renterservices.GetRenterInfo(id)
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
// @Param	token		header		string	true		"The token"
// @Param	body		body 	request.RenterPut	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :renterID is not int
// @router / [put]
func (u *RenterController) Put() {
	id := u.Ctx.Input.Header("rentername")
	var ob request.RenterPut
	err :=json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	err = renterservices.UpdateRenter(id, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title ChangePass
// @Description update the user
// @Param	token		header		string	true		"The token"
// @Param	password	body 	    string	true		"body password"
// @Success 200 {object} models.User
// @Failure 403 :renterID is not int
// @router /password [put]
func (u *RenterController) ChangePass() {
	id := u.Ctx.Input.Header("rentername")
	err := renterservices.ChangePassword(id, string(u.Ctx.Input.RequestBody))
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	token		header 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 renterID is empty
// @router / [delete]
func (u *RenterController) Delete() {
	id := u.Ctx.Input.Header("rentername")
	err := renterservices.DeleteRenter(id)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title GetAllFavoriteHouse
// @Description get all favorite house
// @Param	token		header 	string	true		"The renter favorite house"
// @Success 200 {object} models.House
// @router /favorite [get]
func (u *RenterController) GetAllFavoriteHouse() {
	renterName := u.Ctx.Input.Header("rentername")
	users, err := houseservices.GetAllFavoriteHouse(renterName)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: users,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}