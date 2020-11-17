package housecontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rent-house/models"
	"rent-house/services/commentservices"
	"rent-house/services/houseservices"
)

// Operations about house
type HouseController struct {
	beego.Controller
}

// @Title GetAllActivateHouse
// @Description get all renters
// @Success 200 {object} models.Renter
// @router / [get]
func (u *HouseController) GetAllActivateHouse() {
	users, err := houseservices.GetAllHouse()
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title GetPageActivateHouse
// @Description get page houses
// @Param	page	query	int	true	"page"
// @Param	count	query	int	true	"count"
// @Success 200 {object} models.House
// @router /page [get]
func (u *HouseController) GetPageActivateHouse() {
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
	users, err := houseservices.GetPageHouse(page, count)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	houseID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Renter
// @Failure 403 :houseID is empty
// @router /:houseID/ [get]
func (u *HouseController) Get() {
	id := u.Ctx.Input.Param(":houseID")
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
// @Param	token			header	string	true		"The token string"
// @Param	body		body 	models.Comment	true		"body for user content"
// @Param	houseID	path	string			true		"the house id"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /:houseID/add-comment/ [post]
func (u *HouseController) AddComment() {
	var ob models.Comment
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	houseID := u.Ctx.Input.Param(":houseID")
	err = commentservices.AddComment(houseID, &ob)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	token			header	string	true		"The token string"
// @Param	houseID		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Renter	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :houseID is not int
// @router /:houseID/ [put]
func (u *HouseController) Update() {
	id := u.Ctx.Input.Param(":houseID")
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
// @Param	token			header	string	true		"The token string"
// @Param	houseID		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 houseID is empty
// @router /:houseID/ [delete]
func (u *HouseController) Delete() {
	id := u.GetString(":houseID")
	err := houseservices.DeleteHouse(id)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title GetAllComment
// @Description get all renters
// @Param	houseID	path	string	true	"the house-id
// @Success 200 {object} models.Renter
// @router /:houseID/comments/ [get]
func (u *HouseController) GetAllComment() {
	id := u.Ctx.Input.Param(":houseID")
	users, err := commentservices.GetAllCommentOfHouse(id)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title GetPageComment
// @Description get page comment
// @Param	houseID	path	string	true	"the houseID"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.Comment
// @router /:houseID/page-comments/ [get]
func (u *HouseController) GetPageComment() {
	id := u.Ctx.Input.Param(":houseID")
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
	users, err := commentservices.GetPageCommentOfHouse(id, page, count)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = users
	u.ServeJSON()
}