package houseservices

import (
"rent-house/models"
"rent-house/restapi/response"
)

func AddHouse(house *models.House) error {
	house.Activate = false
	return house.PutItem()
}

func ActiveHouse(id string) error {
	house := &models.House{}
	house, err := house.GetFromKey(id)
	if err != nil {
		return err
	}
	house.Activate = true
	err = house.UpdateItem(id)
	if err != nil {
		return err
	}
	return house.DeleteWaitList(id)
}

func GetHouse(id string) (*models.House, error) {
	o := &models.House{}
	o, err := o.GetFromKey(id)
	return o, err
}

func GetAllWaitHouse() ([]string, error) {
	h := &models.House{}
	return h.GetAllWaitList()
}

func GetPageWaitHouse(page int, count int) ([]string, error) {
	h := &models.House{}
	return h.GetPaginateWaitList(page, count)
}

func GetAllHouse() ([]*response.House, error) {
	o := &models.House{}
	return o.GetAll()
}

func GetAllHouseOfOwner(userID string) ([]*response.House, error) {
	o := &models.House{}
	return o.GetAllHouseOfOwner(userID)
}

func GetPageHouseOfOwner(ownerID string, page int, count int) ([]*response.House, error) {
	o := &models.House{}
	return o.GetPaginateHouseOfUser(ownerID, page, count)
}

func UpdateHouse(id string, ob *models.House) error {
	return ob.UpdateItem(id)
}

func DeleteHouse(id string) error {
	u := &models.House{}
	u, err := u.GetFromKey(id)
	if err != nil {
		return err
	}
	return u.Delete(id)
}