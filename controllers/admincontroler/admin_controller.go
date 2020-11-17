package admincontroler

import (
	"github.com/astaxie/beego"
	"rent-house/services/commentservices"
	"rent-house/services/houseservices"
	"rent-house/services/ownerservices"
)

type AdminController struct {
	beego.Controller
}

// @Title ActivateOwner
// @Description create users
// @Param	key			header	string	true		"admin key"
// @Param	ownerID		query 	string	true		"ownerID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /active-owner/ [post]
func (u *AdminController) ActivateOwner() {
	ownerID := u.GetString("ownerID")
	err := ownerservices.ActiveOwner(ownerID)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title ActivateHouse
// @Description create users
// @Param	key				header	string	true		"admin key"
// @Param	houseID		    query 	string	true		"houseID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /active-house/ [post]
func (u *AdminController) ActivateHouse() {
	houseID := u.GetString("houseID")
	err := houseservices.ActiveHouse(houseID)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title ActivateComment
// @Description create users
// @Param	key				header	string	true		"admin key"
// @Param	commentID		query 	string	true		"houseID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /active-comment/ [post]
func (u *AdminController) ActivateComment() {
	commentID := u.GetString("commentID")
	err := commentservices.ActiveComment(commentID)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "success"
	u.ServeJSON()
}

// @Title GetAllWaitHouse
// @Description get all renters
// @Param	key			header	string	true		"admin key"
// @Success 200 {object} models.House
// @router /wait-houses/ [get]
func (u *AdminController) GetAllWaitHouse() {
	obs, err := houseservices.GetAllWaitHouse()
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = obs
	u.ServeJSON()
}

// @Title GetPageWaitHouse
// @Description get page comment
// @Param	key			header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.House
// @router /page-wait-houses/ [get]
func (u *AdminController) GetPageWaitHouse() {
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
	obs, err := houseservices.GetPageWaitHouse(page, count)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = obs
	u.ServeJSON()
}

// @Title GetAllWaitComment
// @Description get all wait comments
// @Param	key			header	string	true		"admin key"
// @Success 200 {object} models.House
// @router /wait-comments/ [get]
func (u *AdminController) GetAllWaitComment() {
	obs, err := commentservices.GetAllWaitComment()
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = obs
	u.ServeJSON()
}

// @Title GetPageWaitComment
// @Description get page comment
// @Param	key			header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.Comment
// @router /page-wait-comments/ [get]
func (u *AdminController) GetPageWaitComment() {
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
	obs, err := commentservices.GetPageWaitComment(page, count)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = obs
	u.ServeJSON()
}

// @Title GetAllWaitOwner
// @Description get all wait owners
// @Param	key			header	string	true		"admin key"
// @Success 200 {object} models.Owner
// @router /wait-owners/ [get]
func (u *AdminController) GetAllWaitOwner() {
	obs, err := ownerservices.GetAllWaitOwner()
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = obs
	u.ServeJSON()
}

// @Title GetPageWaitOwner
// @Description get page comment
// @Param	key			header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.Owner
// @router /page-wait-owners/ [get]
func (u *AdminController) GetPageWaitOwner() {
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
	obs, err := ownerservices.GetPageWaitOwner(page, count)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = obs
	u.ServeJSON()
}

