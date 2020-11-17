package houseservices

import (
	"mime/multipart"
	"rent-house/models"
"rent-house/restapi/response"
)

func AddHouse(ownerID string, house *models.House) (string, error) {
	house.OwnerID = ownerID
	house.Activate = false
	return house.PutItem()
}

func ActiveHouse(id string) error {
	house := &models.House{}
	err := house.GetFromKey(id)
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

func UploadFile(houseID string, file []*multipart.FileHeader) error {
	house := &models.House{}
	err := house.GetFromKey(houseID)
	if err != nil {
		return err
	}
	for _, i := range file {
		f, err := i.Open()
		if err != nil {
			return err
		}
		err = house.AddImage(f,houseID)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetHouse(id string) (*models.House, error) {
	o := &models.House{}
	err := o.GetFromKey(id)
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
	return o.GetAllActivate()
}

func GetPageHouse(page, count int) ([]*response.House, error) {
	o := &models.House{}
	return o.GetPageActivate(page, count)
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
	err := u.GetFromKey(id)
	if err != nil {
		return err
	}
	return u.Delete(id)
}