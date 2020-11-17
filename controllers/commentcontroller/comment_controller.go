package commentcontroller

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rent-house/models"
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
	if id != "" {
		comment, err := commentservices.GetComment(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = comment
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	token			header	string	true		"The token string"
// @Param	commentID		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Comment	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :commentID is not int
// @router /:commentID/ [put]
func (u *CommentController) Update() {
	id := u.Ctx.Input.Param(":commentID")
	if id != "" {
		var ob models.Comment
		err :=json.Unmarshal(u.Ctx.Input.RequestBody, &ob)
		if err != nil {
			u.Ctx.WriteString(err.Error())
			return
		}
		err = commentservices.UpdateComment(id, &ob)
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
// @Param	commentID		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 commentID is empty
// @router /:commentID/ [delete]
func (u *CommentController) Delete() {
	id := u.GetString(":commentID")
	err := commentservices.DeleteComment(id)
	if err != nil {
		u.Ctx.WriteString(err.Error())
		return
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
