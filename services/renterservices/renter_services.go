package renterservices

import (
	"errors"
	"rent-house/middlewares"
	"rent-house/models"
	"rent-house/restapi/response"
)

func AddRenter(o *models.Renter) error {
	_, err := o.GetFromKey(o.RenterName)
	if err != nil {
		err = o.PutItem()
	}
	return errors.New("already exist")
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

func LoginRenter(login models.RenterLogin) (string, error) {
	renter := &models.Renter{}
	renter, err := renter.GetFromKey(login.RenterName)
	if err != nil {
		return "", err
	}
	if login.Password == renter.Password {
		return middlewares.CreateToken(login)
	}
	return "", errors.New("not authorized")
}
