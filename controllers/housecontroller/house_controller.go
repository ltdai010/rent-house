package housecontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"rent-house/services/commentservices"
	"rent-house/services/houseservices"
)

// Operations about house
type HouseController struct {
	beego.Controller
}

// @Title GetAllActivateHouse
// @Description get all renters
// @Success 200 {object} models.House
// @router / [get]
func (u *HouseController) GetAllActivateHouse() {
	users, err := houseservices.GetAllActiveHouse()
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

// @Title GetPageActivateHouse
// @Description get page houses
// @Param	page	query	int	true	"page"
// @Param	count	query	int	true	"count"
// @Success 200 {object} models.House
// @router /page [get]
func (u *HouseController) GetPageActivateHouse() {
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
	users, total, err := houseservices.GetPageActiveHouse(page, count)
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonArray{
			Data: users,
			TotalCount: int64(total),
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	houseID		path 	string	true		"The house id"
// @Success 200 {object} models.House
// @Failure 403 :houseID is empty
// @router /:houseID/ [get]
func (u *HouseController) Get() {
	id := u.Ctx.Input.Param(":houseID")
	//get house
	house, err := houseservices.GetHouse(id)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	//raise view
	go houseservices.ViewHouse(id)
	u.Data["json"] = response.ResponseCommonSingle{
		Data: house,
		Err:  response.NewErr(response.Success),
	}
	u.ServeJSON()
}



// @Title GetByLike
// @Description get user by uid
// @Param	page		query 	int	false		"The house id"
// @Param	count		query	int	false		"page length"
// @Success 200 {object} models.House
// @Failure 403 :houseID is empty
// @router /favorite-desc [get]
func (u *HouseController) GetByLike(page, count int) {
	//get house
	house, total, err := houseservices.GetHouseArrangeByLike(page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	u.Data["json"] = response.ResponseCommonArray{
		Data: house,
		TotalCount: int64(total),
		Err:  response.NewErr(response.Success),
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the house
// @Param	token		header	string	true		"The token string"
// @Param	houseID		path 	string	true		"The uid you want to update"
// @Param	body		body 	request.HousePut	true		"body for user content"
// @Success 200 {string} success
// @Failure 403 :houseID is not int
// @router /:houseID/ [put]
func (u *HouseController) Update() {
	id := u.Ctx.Input.Param(":houseID")
	var ob request.HousePut
	err :=json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	err = houseservices.UpdateHouse(id, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title UpdateRented
// @Description update the house
// @Param	token		header	string	true		"The token string"
// @Param	houseID		path 	string	true		"The uid you want to update"
// @Success 200 {string} success
// @Failure 403 :houseID is not int
// @router /:houseID/rented [put]
func (u *HouseController) UpdateRented() {
	id := u.Ctx.Input.Param(":houseID")
	err := houseservices.UpdateRented(id)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}


// @Title UpdateExpiredTime
// @Description update the user
// @Param	token		header	string	true		"The token string"
// @Param	houseID		path 	string	true		"The uid you want to update"
// @Param	time		query 	int64	true		"time"
// @Success 200 {string} success
// @Failure 403 :houseID is not int
// @router /:houseID/extend-house [put]
func (u *HouseController) UpdateExpiredTime(time int64) {
	id := u.Ctx.Input.Param(":houseID")
	err := houseservices.PutExtendTime(id, time)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title UploadImage
// @Description create users
// @Param	files		formData    []file			true		"house image"
// @Success 200 {[]string} image link
// @Failure 403 body is empty
// @router /images [post]
func (u *HouseController) UploadImage() {
	file, err := u.GetFiles("files")
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	list, err := houseservices.UploadFile(file)
	if err != nil {
		u.Data["json"] = response.NewErr(response.ErrSystem)
		u.ServeJSON()
		return
	}
	u.Data["json"] = response.ResponseCommonSingle{
		Data: list,
		Err:  response.NewErr(response.Success),
	}
	u.ServeJSON()
	return
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
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title GetAllComment
// @Description get all renters
// @Param	houseID	path	string	true	"the house-id
// @Success 200 {object} response.Comment
// @router /:houseID/comments/ [get]
func (u *HouseController) GetAllComment() {
	id := u.Ctx.Input.Param(":houseID")
	comments, err := commentservices.GetAllCommentActiveOfHouse(id)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: comments,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageComment
// @Description get page comment
// @Param	houseID	path	string	true	"the houseID"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} response.Comment
// @router /:houseID/page-comments/ [get]
func (u *HouseController) GetPageComment() {
	id := u.Ctx.Input.Param(":houseID")
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
	users, total, err := commentservices.GetPageActiveCommentOfHouse(id, page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonArray{
			Data: users,
			TotalCount: int64(total),
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}