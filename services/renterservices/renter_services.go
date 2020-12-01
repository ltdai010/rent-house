package renterservices

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"rent-house/middlewares"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
)

func AddRenter(o *request.RenterPost) error {
	r := &models.Renter{}
	err := r.GetFromKey(o.RenterName)
	if err != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(o.Password), bcrypt.DefaultCost)
		if err != nil {
			return response.ErrSystem
		}
		r = &models.Renter{
			RenterName:     o.RenterName,
			RenterFullName: o.RenterFullName,
			Password:       string(hashed),
			PhoneNumber:    o.PhoneNumber,
			Email:          o.Email,
			ListFavourite:  []string{},
		}
		err = r.PutItem()
		return err
	}
	return errors.New("already exist")
}

func GetRenter(id string) (models.Renter, error) {
	o := &models.Renter{}
	err := o.GetFromKey(id)
	if err != nil {
		return models.Renter{}, err
	}
	return *o, err
}

func GetAllRenter() ([]response.Renter, error) {
	o := &models.Renter{}
	list, err := o.GetAll()
	if err != nil {
		return []response.Renter{}, err
	}
	return list, nil
}

func UpdateRenter(id string, ob *request.RenterPut) error {
	o := &models.Renter{}
	err := o.GetFromKey(id)
	if err != nil {
		return err
	}
	o.Password = ob.Password
	o.Email = ob.Email
	o.PhoneNumber = ob.PhoneNumber
	o.RenterFullName = ob.RenterFullName
	return o.UpdateItem(id)
}

func DeleteRenter(id string) error {
	u := &models.Renter{}
	err := u.GetFromKey(id)
	if err != nil {
		return err
	}
	return u.Delete(id)
}

func LoginRenter(login models.Login) (string, error) {
	renter := &models.Renter{}
	err := renter.GetFromKey(login.Username)
	if err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(renter.Password), []byte(login.Password)) == nil {
		return middlewares.CreateToken(login.Username)
	}
	return "", errors.New("not authorized")
}


