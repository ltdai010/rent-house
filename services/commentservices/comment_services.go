package commentservices

import (
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"time"
)

func AddComment(houseID string, ownerID string, ob *request.CommentPost) error {
	c := &models.Comment{
		Content:  ob.Content,
		OwnerID:  ownerID,
		Header:   ob.Header,
		HouseID:  houseID,
		PostTime: time.Now().Unix(),
		Star:     0,
		Activate: false,
	}
	return c.PutItem()
}

func ActiveComment(id string) error {
	comment := &models.Comment{}
	err := comment.GetFromKey(id)
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
	list, err := h.GetAllWaitList()
	if err != nil {
		return []string{}, err
	}
	return list, nil
}

func GetPageWaitComment(page int, count int) ([]string, error) {
	h := &models.Comment{}
	list, err := h.GetPaginateWaitList(page, count)
	if err != nil {
		return []string{}, err
	}
	return list, nil
}

func GetComment(id string) (models.Comment, error) {
	o := &models.Comment{}
	err := o.GetFromKey(id)
	if err != nil {
		return models.Comment{}, err
	}
	return *o, nil
}

func GetAllComment() ([]response.Comment, error) {
	o := &models.Comment{}
	list, err := o.GetAll()
	if err != nil {
		return []response.Comment{}, err
	}
	return list, nil
}

func GetAllCommentOfHouse(houseID string) ([]response.Comment, error) {
	o := &models.Comment{}
	list, err := o.GetAllCommentInPost(houseID)
	if err != nil {
		return []response.Comment{}, err
	}
	return list, nil
}

func GetPageCommentOfHouse(houseID string, page int, count int) ([]response.Comment, error) {
	o := &models.Comment{}
	list, err := o.GetPaginateCommentInHouse(houseID, page, count)
	if err != nil {
		return []response.Comment{}, err
	}
	return list, nil
}

func UpdateComment(id string, ob *request.CommentPut) error {
	c := &models.Comment{}
	err := c.GetFromKey(id)
	if err != nil {
		return err
	}
	c.Header = ob.Header
	c.Content = ob.Content
	return c.UpdateItem(id)
}

func DeleteComment(id string) error {
	u := &models.Comment{}
	err := u.GetFromKey(id)
	if err != nil {
		return err
	}
	return u.Delete(id)
}
