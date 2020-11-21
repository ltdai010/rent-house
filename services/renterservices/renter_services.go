package renterservices

import (
	"errors"
	"rent-house/middlewares"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
)

func AddRenter(o *request.RenterPost) error {
	r := &models.Renter{}
	err := r.GetFromKey(o.RenterName)
	if err != nil {
		err = r.PutItem()
		return nil
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
	if login.Password == renter.Password {
		return middlewares.CreateToken(login.Username)
	}
	return "", errors.New("not authorized")
}

func AddToFavourite(renterID, houseID string) error {
	r := &models.Renter{}
	err := r.GetFromKey(renterID)
	if err != nil {
		return err
	}
	//check if already exist
	for _, i := range r.ListFavourite {
		if i == houseID {
			return response.Existed
		}
	}
	//add to list
	r.ListFavourite = append(r.ListFavourite, houseID)
	err = r.PutItem()
	if err != nil {
		return err
	}
	//raise the number of likes
	h := &models.House{}
	err = h.GetFromKey(houseID)
	if err != nil {
		return err
	}
	h.Like++
	_, err = h.PutItem()
	return nil
}
