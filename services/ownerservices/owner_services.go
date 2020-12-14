package ownerservices

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"rent-house/middlewares"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"time"
)

func AddOwner(o *request.OwnerPost) error {
	ob := &models.Owner{}
	err := ob.GetFromKey(o.OwnerName)
	if err == nil {
		return errors.New("already exist")
	}
	ad := &models.Admin{}
	err = ad.GetFromKey(o.OwnerName)
	if err == nil {
		return errors.New("already exist")
	}
	a := &models.Address{}
	err = a.FindAddress(o.CommuneCode)
	if err != nil {
		log.Println(err)
		return err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(o.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrSystem
	}
	ob = &models.Owner{
		OwnerName:       o.OwnerName,
		Password:        string(hashed),
		OwnerFullName:   o.OwnerFullName,
		Profile:         o.Profile,
		Address:         *a,
		Activate:        false,
		PostTime:        time.Now().Unix(),
		PasswordChanged: time.Now().Unix(),
	}
	err = ob.PutItem()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func ActiveOwner(ownerID string) error {
	owner := &models.Owner{}
	err := owner.GetFromKey(ownerID)
	if err != nil {
		return err
	}
	owner.Activate = true
	owner.PostTime = time.Now().Unix()
	err = owner.UpdateItem(ownerID)
	if err != nil {
		return err
	}
	return nil
}

func DeactiveOwner(ownerID string) error {
	owner := &models.Owner{}
	err := owner.GetFromKey(ownerID)
	if err != nil {
		return err
	}
	owner.Activate = false
	owner.PostTime = time.Now().Unix()
	err = owner.UpdateItem(ownerID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllWaitOwner() ([]response.Owner, error) {
	h := &models.Owner{}
	list, err := h.GetAllWaitList()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetPageWaitOwner(page int, count int) ([]response.Owner, error) {
	h := &models.Owner{}
	list, err := h.GetPaginateWaitList(page, count)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetOwner(ownerID string) (response.Owner, error) {
	u := &models.Owner{}
	err := u.GetFromKey(ownerID)
	if err != nil {
		return response.Owner{}, err
	}
	res := response.Owner{
		OwnerName:   ownerID,
		OwnerFullName: u.OwnerFullName,
		Profile:   response.Profile{
			IDCard:      u.Profile.IDCard,
			PhoneNumber: u.Profile.PhoneNumber,
			Email:       u.Profile.Email,
		},
		Address:   response.Address{
			Province: u.Address.Province,
			District: u.Address.District,
			Commune:  u.Address.Commune,
		},
		AverageStar: AverageStar(ownerID),
		Activate:  u.Activate,
	}
	return res, nil
}

func AverageStar(ownerID string) float32 {
	house := &models.House{}
	list, err := house.GetAllHouseOfOwner(ownerID)
	if err != nil {
		return 5
	}
	var num float32
	var sum float32
	for _, i := range list {
		if i.Review == nil {
			continue
		}
		n := i.Review["0"] + i.Review["1"] + i.Review["2"] + i.Review["3"] + i.Review["4"] + i.Review["5"]
		s := i.Review["1"]*1 + i.Review["2"]*2 + i.Review["3"]*3 + i.Review["4"]*4 + i.Review["5"]*5
		if n == 0 {
			continue
		}
		sum += float32(s)/float32(n)
		num++
	}
	if num == 0 {
		return 5
	}
	return sum / num
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
	if o.Activate == true {
		return response.NotPermission
	}
	a := &models.Address{}
	err = a.FindAddress(ob.CommuneCode)
	if err != nil {
		return err
	}
	o.OwnerFullName = ob.OwnerFullName
	o.Address = *a
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
	if bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(login.Password)) == nil {
		return middlewares.CreateToken(login.Username)
	}
	return "", errors.New("not authorized")
}

func ChangePassword(ownerID string, password string) error {
	owner := &models.Owner{}
	err := owner.GetFromKey(ownerID)
	if err != nil {
		return err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	owner.PasswordChanged = time.Now().Unix()
	owner.Password = string(hashed)
	return owner.PutItem()
}