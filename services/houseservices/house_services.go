package houseservices

import (
"rent-house/models"
"rent-house/restapi/response"
)

func AddHouse(house *models.House) error {
	return house.PutItem()
}

func GetHouse(id string) (*models.House, error) {
	o := &models.House{}
	o, err := o.GetFromKey(id)
	return o, err
}

func GetAllHouse() ([]*response.House, error) {
	o := &models.House{}
	return o.GetAll()
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