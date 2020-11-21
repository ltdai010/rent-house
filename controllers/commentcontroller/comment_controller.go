package commentcontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"rent-house/services/commentservices"
)

type CommentController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	token			header	string	true		"The token string"
// @Param	commentID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comment
// @Failure 403 :commentID is empty
// @router /:commentID/ [get]
func (u *CommentController) Get() {
	id := u.Ctx.Input.Param(":commentID")
	comment, err := commentservices.GetComment(id)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: comment,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	token			header	string	true		"The token string"
// @Param	commentID		path 	string	true		"The uid you want to update"
// @Param	body		body 	 request.CommentPut	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :commentID is not int
// @router /:commentID/ [put]
func (u *CommentController) Update() {
	id := u.Ctx.Input.Param(":commentID")
	var ob request.CommentPut
	err :=json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	err = commentservices.UpdateComment(id, &ob)
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
// @Param	commentID		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 commentID is empty
// @router /:commentID/ [delete]
func (u *CommentController) Delete() {
	id := u.GetString(":commentID")
	err := commentservices.DeleteComment(id)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.NewErr(response.Success)
	}
	u.ServeJSON()
}
