package renterservices

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func AddRenter(o *models.Renter) error {
	return o.PutItem()
}

func GetRenter(id string) (*models.Renter, error) {
	o := &models.Renter{}
	o, err := o.GetFromKey(id)
	return o, err
}

func GetAllRenter() ([]*response.Renter, error) {
	o := &models.Renter{}
	return o.GetAll()
}

func UpdateRenter(id string, ob *models.Renter) error {
	return ob.UpdateItem(id)
}

func DeleteRenter(id string) error {
	u := &models.Renter{}
	u, err := u.GetFromKey(id)
	if err != nil {
		return err
	}
	return u.Delete(id)
}
