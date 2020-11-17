package ownerservices

import (
	"errors"
	"rent-house/middlewares"
	"rent-house/models"
	"rent-house/restapi/response"
)

func AddOwner(o *models.Owner) error {
	o.Activate = false
	return o.PutItem()
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
	return h.GetAllWaitList()
}

func GetPageWaitOwner(page int, count int) ([]string, error) {
	h := &models.Owner{}
	return h.GetPaginateWaitList(page, count)
}

func GetOwner(ownerID string) (*models.Owner, error) {
	u := &models.Owner{}
	err := u.GetFromKey(ownerID)
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