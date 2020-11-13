package commentservices

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func AddComment(ob *models.Comment) error {
	ob.Activate = false
	return ob.PutItem()
}

func ActiveComment(id string) error {
	comment := &models.Comment{}
	comment, err := comment.GetFromKey(id)
	if err != nil {
		return err
	}
	comment.Activate = true
	err = comment.UpdateItem(id)
	if err != nil {
		return err
	}
	return comment.DeleteWaitList(id)
}

func GetAllWaitComment() ([]string, error) {
	h := &models.Comment{}
	return h.GetAllWaitList()
}

func GetPageWaitComment(page int, count int) ([]string, error) {
	h := &models.Comment{}
	return h.GetPaginateWaitList(page, count)
}

func GetComment(id string) (*models.Comment, error) {
	o := &models.Comment{}
	o, err := o.GetFromKey(id)
	return o, err
}

func GetAllComment() ([]*response.Comment, error) {
	o := &models.Comment{}
	return o.GetAll()
}

func GetAllCommentOfHouse(houseID string) ([]*response.Comment, error) {
	o := &models.Comment{}
	return o.GetAllCommentInPost(houseID)
}

func GetPageCommentOfHouse(houseID string, page int, count int) ([]*response.Comment, error) {
	o := &models.Comment{}
	return o.GetPaginateCommentInPost(houseID, page, count)
}

func UpdateComment(id string, ob *models.Comment) error {
	return ob.UpdateItem(id)
}

func DeleteComment(id string) error {
	u := &models.Comment{}
	u, err := u.GetFromKey(id)
	if err != nil {
		return err
	}
	return u.Delete(id)
}
