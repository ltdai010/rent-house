package ownerservices

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func AddOwner(o *models.Owner) error {
	return o.PutItem()
}

func GetOwner(ownerID string) (*models.Owner, error) {
	u := &models.Owner{}
	u, err := u.GetFromKey(ownerID)
	return u, err
}

func GetAllOwner() ([]*response.Owner, error) {
	u := &models.Owner{}
	return u.GetAll()
}

func UpdateUser(id string, o *models.Owner) error {
	return o.UpdateItem(id)
}

func DeleteOwner(ownerID string) error {
	u := &models.Owner{}
	u, err := u.GetFromKey(ownerID)
	if err != nil {
		return err
	}
	return u.Delete(ownerID)
}