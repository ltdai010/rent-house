package commentservices

import (
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"strconv"
	"time"
)

func AddComment(houseID string, ownerID string, ob *request.CommentPost) error {
	if ob.Star < 0 || ob.Star > 5 {
		return response.BadRequest
	}
	c := &models.Comment{
		Content:  ob.Content,
		RenterID: ownerID,
		HouseID:  houseID,
		PostTime: time.Now().Unix(),
		Star:     ob.Star,
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
	//update house
	h := &models.House{}
	err = h.GetFromKey(comment.HouseID)
	if err != nil {
		return err
	}
	if h.Review == nil {
		h.Review = map[string]int{}
	}
	if v, ok := h.Review[strconv.Itoa(comment.Star)]; ok {
		h.Review[strconv.Itoa(comment.Star)] = v + 1
	} else {
		h.Review[strconv.Itoa(comment.Star)] = 1
	}
	//update
	go h.UpdateItem(comment.HouseID)
	return nil
}

func GetAllWaitComment() ([]response.Comment, error) {
	h := &models.Comment{}
	list, err := h.GetAllWaitList()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetPageWaitComment(page int, count int) ([]response.Comment, int, error) {
	h := &models.Comment{}
	return h.GetPaginateWaitList(page, count)
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

func GetAllCommentActiveOfHouse(houseID string) ([]response.Comment, error) {
	o := &models.Comment{}
	list, err := o.GetAllCommentActiveInHouse(houseID)
	if err != nil {
		return []response.Comment{}, err
	}
	return list, nil
}

func GetPageActiveCommentOfHouse(houseID string, page int, count int) ([]response.Comment, int, error) {
	o := &models.Comment{}
	return o.GetPaginateCommentActiveInHouse(houseID, page, count)
}

func GetAllCommentOfHouse(houseID string) ([]response.Comment, error) {
	o := &models.Comment{}
	list, err := o.GetAllCommentInHouse(houseID)
	if err != nil {
		return []response.Comment{}, err
	}
	return list, nil
}

func GetPageCommentOfHouse(houseID string, page int, count int) ([]response.Comment, int, error) {
	o := &models.Comment{}
	return o.GetPaginateCommentInHouse(houseID, page, count)
}

func UpdateComment(id string, ob *request.CommentPut) error {
	c := &models.Comment{}
	err := c.GetFromKey(id)
	if err != nil {
		return err
	}
	if c.Activate {
		return response.NotPermission
	}
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
