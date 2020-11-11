package ownerservices

import (
	"fmt"
	"rent-house/consts"
	"rent-house/models"
	"rent-house/restapi/request"
)

func AddUser(user *request.PostOwner) error {
	if err != nil {
		return err
	}
	u := &models.User{
		UserID:      fmt.Sprintf(consts.ID_FORMAT, i),
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Address:     user.Address,
	}
	err = u.PutItem(u.UserID)
	return err
}

func GetUser(userID string) (*models.User, error) {
	u := &models.User{}
	u, err := u.GetFromKey(userID)
	return u, err
}

func GetAllUser() ([]models.User, int64, error) {
	u := &models.User{}
	list, leng, err := u.GetAll()
	return list, leng, err
}

func UpdateUser(user *request.PutUser) error {
	u := &models.User{}
	u, err := u.GetFromKey(user.UserID)
	if err != nil {
		return err
	}
	u.Username = user.Username
	u.Address = user.Address
	u.Email = user.Email
	u.PhoneNumber = user.PhoneNumber
	return u.PutItem(user.UserID)
}

func DeleteUser(userID string) error {
	u := &models.User{}
	u, err := u.GetFromKey(userID)
	if err != nil {
		return err
	}
	return u.Delete()
}