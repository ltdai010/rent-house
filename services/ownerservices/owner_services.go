package ownerservices

import (
	"errors"
	"rent-house/middlewares"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
)

func AddOwner(o *request.OwnerPost) error {
	ob := &models.Owner{}
	err := ob.GetFromKey(o.OwnerFullName)
	if err == nil {
		return errors.New("already exist")
	}
	ob = &models.Owner{
		OwnerName:     o.OwnerName,
		Password:      o.Password,
		OwnerFullName: o.OwnerFullName,
		Profile:       o.Profile,
		Address:       o.Address,
		Activate:      false,
	}
	return ob.PutItem()
}

func ActiveOwner(ownerID string) error {
	owner := &models.Owner{}
	err := owner.GetFromKey(ownerID)
	if err != nil {
		return err
	}
	owner.Activate = true
	err = owner.UpdateItem(ownerID)
	if err != nil {
		return err
	}
	return owner.DeleteWaitList(ownerID)
}

func GetAllWaitOwner() ([]string, error) {
	h := &models.Owner{}
	list, err := h.GetAllWaitList()
	if err != nil {
		return []string{}, err
	}
	return list, nil
}

func GetPageWaitOwner(page int, count int) ([]string, error) {
	h := &models.Owner{}
	list, err := h.GetPaginateWaitList(page, count)
	if err != nil {
		return []string{}, err
	}
	return list, nil
}

func GetOwner(ownerID string) (models.Owner, error) {
	u := &models.Owner{}
	err := u.GetFromKey(ownerID)
	if err != nil {
		return models.Owner{}, err
	}
	return *u, nil
}

func GetAllOwner() ([]response.Owner, error) {
	u := &models.Owner{}
	list, err := u.GetAll()
	if err != nil {
		return []response.Owner{}, err
	}
	return list, nil
}

func GetPageOwner(page, count int) ([]response.Owner, error) {
	u := &models.Owner{}
	list, err := u.GetPaginate(page, count)
	if err != nil {
		return []response.Owner{}, err
	}
	return list, nil
}

func UpdateOwner(id string, ob *request.OwnerPut) error {
	o := &models.Owner{}
	err := o.GetFromKey(id)
	if err != nil {
		return err
	}
	o.Password = ob.Password
	o.OwnerFullName = ob.OwnerFullName
	o.Address = ob.Address
	o.Profile = ob.Profile
	return o.UpdateItem(id)
}

func DeleteOwner(ownerID string) error {
	u := &models.Owner{}
	err := u.GetFromKey(ownerID)
	if err != nil {
		return err
	}
	return u.Delete(ownerID)
}

func LoginOwner(login models.Login) (string, error) {
	owner := &models.Owner{}
	err := owner.GetFromKey(login.Username)
	if err != nil {
		return "", err
	}
	if login.Password == owner.Password {
		return middlewares.CreateToken(login.Username)
	}
	return "", errors.New("not authorized")
}