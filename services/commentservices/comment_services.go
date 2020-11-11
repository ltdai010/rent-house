package commentservices

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func AddComment(ob *models.Comment) error {
	return ob.PutItem()
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
