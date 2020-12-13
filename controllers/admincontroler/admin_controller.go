package admincontroler

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"rent-house/services/adminservice"
	"rent-house/services/commentservices"
	"rent-house/services/houseservices"
	"rent-house/services/ownerservices"
	"rent-house/services/renterservices"
)

type AdminController struct {
	beego.Controller
}

// @Title ActivateOwner
// @Description create users
// @Param	token		header	string	true		"admin key"
// @Param	ownerID		query 	string	true		"ownerID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /active-owner/ [post]
func (u *AdminController) ActivateOwner() {
	ownerID := u.GetString("ownerID")
	err := ownerservices.ActiveOwner(ownerID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title CreateHouse
// @Description create users month = 0|| quarter = 1|| year = 2
// @Param	token		header	    string			true			"The token string"
// @Param	body		body 		request.HousePost	true		"body for user content"
// @Success 200 {int} models.House
// @Failure 403 body is empty
// @router /house/ [post]
func (u *AdminController) CreateHouse() {
	var ob request.HousePost
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		log.Println(err)
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	ownerID := u.Ctx.Input.Header("admin")
	s, err := houseservices.AdminAddHouse(ownerID, &ob)
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


// @Title Login
// @Description login
// @Param	login		body 	models.Login	true		"body for user content"
// @Success 200 {string} token
// @Failure 403 body is empty
// @router /login/ [post]
func (u *AdminController) Login() {
	var ob models.Login
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	token, err := adminservice.LoginAdmin(ob)
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

// @Title ActivateHouse
// @Description active house
// @Param	token			header	string	true		"admin key"
// @Param	houseID		    query 	string	true		"houseID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /active-house/ [post]
func (u *AdminController) ActivateHouse() {
	houseID := u.GetString("houseID")
	err := houseservices.ActiveHouse(houseID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title DeniedHouse
// @Description active house
// @Param	token			header	string	true		"admin key"
// @Param	houseID		    query 	string	true		"houseID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /denied-house/ [post]
func (u *AdminController) DeniedHouse() {
	houseID := u.GetString("houseID")
	err := houseservices.DenyHouse(houseID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title ExtendHouse
// @Description extend house time
// @Param	token			header	string	true		"admin key"
// @Param	houseID		    query 	string	true		"houseID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /extend-house/ [post]
func (u *AdminController) ExtendHouse() {
	houseID := u.GetString("houseID")
	err := houseservices.ExtendHouseTime(houseID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title ActivateComment
// @Description create users
// @Param	token			header	string	true		"admin key"
// @Param	commentID		query 	string	true		"houseID"
// @Success 200 {string} success
// @Failure 403 body is empty
// @router /active-comment/ [post]
func (u *AdminController) ActivateComment() {
	commentID := u.GetString("commentID")
	err := commentservices.ActiveComment(commentID)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}

// @Title GetAllWaitHouse
// @Description get all renters
// @Param	token			header	string	true		"admin key"
// @Success 200 {object} models.House
// @router /wait-houses/ [get]
func (u *AdminController) GetAllWaitHouse() {
	obs, err := houseservices.GetAllHouseHouseByStatus(models.InActivated)
	if err != nil {
		u.Data["json"] = response.NewErr(response.ErrSystem)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageWaitHouse
// @Description get page comment
// @Param	token		header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.House
// @router /page-wait-houses/ [get]
func (u *AdminController) GetPageWaitHouse() {
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
	obs, err := houseservices.GetPageHouseByStatus(models.InActivated, page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}


// @Title GetAllExtendHouse
// @Description get all renters
// @Param	token			header	string	true		"admin key"
// @Success 200 {object} models.House
// @router /extend-houses/ [get]
func (u *AdminController) GetAllExtendHouse() {
	obs, err := houseservices.GetAllHouseHouseByStatus(models.Extend)
	if err != nil {
		u.Data["json"] = response.NewErr(response.ErrSystem)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageExtendHouse
// @Description get page comment
// @Param	token		header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.House
// @router /page-extend-houses/ [get]
func (u *AdminController) GetPageExtendHouse() {
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
	obs, err := houseservices.GetPageHouseByStatus(models.Extend, page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}


// @Title GetAllDeniedHouse
// @Description get all renters
// @Param	token			header	string	true		"admin key"
// @Success 200 {object} models.House
// @router /denied-houses/ [get]
func (u *AdminController) GetAllDeniedHouse() {
	obs, err := houseservices.GetAllHouseHouseByStatus(models.Denied)
	if err != nil {
		u.Data["json"] = response.NewErr(response.ErrSystem)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageDeniedHouse
// @Description get page comment
// @Param	token		header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.House
// @router /page-denied-houses/ [get]
func (u *AdminController) GetPageDeniedHouse() {
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
	obs, err := houseservices.GetPageHouseByStatus(models.Denied, page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetAllWaitComment
// @Description get all wait comments
// @Param	token			header	string	true		"admin key"
// @Success 200 {object} models.House
// @router /wait-comments/ [get]
func (u *AdminController) GetAllWaitComment() {
	obs, err := commentservices.GetAllWaitComment()
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageWaitComment
// @Description get page comment
// @Param	token			header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.Comment
// @router /page-wait-comments/ [get]
func (u *AdminController) GetPageWaitComment() {
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
	obs, err := commentservices.GetPageWaitComment(page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetAllOwner
// @Description get all owners
// @Param	token			header	string	true		"admin key"
// @Success 200 {object} models.Owner
// @router /owners/ [get]
func (u *AdminController) GetAllOwner() {
	users, err := ownerservices.GetAllOwner()
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

// @Title GetAllRenter
// @Description get all renters
// @Param	token			header	string	true		"admin key"
// @Success 200 {object} models.Renter
// @router /renters/ [get]
func (u *AdminController) GetAllRenter() {
	users, err := renterservices.GetAllRenter()
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

// @Title GetAllWaitOwner
// @Description get all wait owners
// @Param	token			header	string	true		"admin key"
// @Success 200 {object} models.Owner
// @router /wait-owners/ [get]
func (u *AdminController) GetAllWaitOwner() {
	obs, err := ownerservices.GetAllWaitOwner()
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageWaitOwner
// @Description get page comment
// @Param	token			header	string	true		"admin key"
// @Param	page		query	int		true	"the page"
// @Param	count		query	int		true	"the count"
// @Success 200 {object} models.Owner
// @router /page-wait-owners/ [get]
func (u *AdminController) GetPageWaitOwner() {
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
	obs, err := ownerservices.GetPageWaitOwner(page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: obs,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

