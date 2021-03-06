package ownercontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"rent-house/services/houseservices"
	"rent-house/services/ownerservices"
	services2 "rent-house/websocket/chatservice/services"
	"rent-house/websocket/notificationservice/services"
)

// Operations about Owner
type OwnerController struct {
	beego.Controller
}

// @Title CreateHouse
// @Description create users month = 0|| quarter = 1|| year = 2
// @Param	token		header	    string			true		"The token string"
// @Param	body		body 		request.HousePost	true		"body for user content"
// @Success 200 {int} models.House
// @Failure 403 body is empty
// @router /house/ [post]
func (u *OwnerController) CreateHouse() {
	var ob request.HousePost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	ownerID := u.Ctx.Input.Header("ownername")
	s, err := houseservices.AddHouse(ownerID, &ob)
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	u.Data["json"] = response.ResponseCommonSingle{
		Data: s,
		Err:  response.NewErr(response.Success),
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
func (u *OwnerController) ChangePass() {
	id := u.Ctx.Input.Header("ownername")
	err := ownerservices.ChangePassword(id, string(u.Ctx.Input.RequestBody))
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
// @Param   ownerID	path   string	true		"The ownerID"
// @Success 200 {object} models.User
// @Failure 403 :ownerID is empty
// @router /:ownerID [get]
func (u *OwnerController) GetOwner(ownerID string) {
	user, err := ownerservices.GetOwner(ownerID)
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

// @Title GetMessage
// @Description get user by uid
// @Param	token		header	    string			true		"The token string"
// @Param	page		query		int				false		"the page number"
// @Param	length		query		int				false		"the page  length"
// @Success 200 {object} models.ResMessageConversation
// @Failure 403 :ownerID is empty
// @router /messages/ [get]
func (u *OwnerController) GetMessage(page, length int) {
	ownername := u.Ctx.Input.Header("ownername")
	user, total, err := services2.GetMessageOfOwner(ownername, page, length)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonArray{
			Data: user,
			TotalCount: int64(total),
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	token			header	string	true		"The token string"
// @Param	body		body 	request.OwnerPut	true		"body for user content"
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
// @Param	ownerID			path	string	true		"The ownerID string"
// @Success 200 {object} models.House
// @router /:ownerID/houses/ [get]
func (u *OwnerController) GetAllHouse(ownerID string) {
	houses, err := houseservices.GetAllHouseOfOwner(ownerID)
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

// @Title GetNotice
// @Description get all renters
// @Param	token		header	string	true		"The owner name string"
// @Param	page		query	int	false			"The page"
// @Param	length		query	int	false			"Page length"
// @Success 200 {object} models.ResNotification
// @router /notification/ [get]
func (u *OwnerController) GetAllNotice(page, length int) {
	ownername := u.Ctx.Input.Header("ownername")
	if length > 0 {
		notes, total, err := services.GetPageNotificationOfOwner(ownername, page, length)
		if err != nil {
			log.Println(err)
			u.Data["json"] = response.NewErr(response.BadRequest)
		} else {
			u.Data["json"] = response.ResponseCommonArray{
				Data: notes,
				TotalCount: int64(total),
				Err:  response.NewErr(response.Success),
			}
		}
	} else {
		notes, err := services.GetAllNotificationOfOwner(ownername)
		if err != nil {
			u.Data["json"] = response.NewErr(response.BadRequest)
		} else {
			u.Data["json"] = response.ResponseCommonSingle{
				Data: notes,
				Err:  response.NewErr(response.Success),
			}
		}
	}
	u.ServeJSON()
}

// @Title GetPageHouse
// @Description get page house
// @Param	ownerID		path	string	true		"The ownerID "
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.House
// @router /:ownerID/page-houses/ [get]
func (u *OwnerController) GetPageHouse(ownerID string, page int, count int) {
	houses, total, err := houseservices.GetPageHouseOfOwner(ownerID, page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonArray{
			Data: houses,
			TotalCount: int64(total),
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