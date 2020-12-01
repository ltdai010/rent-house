package houseservices

import (
	"log"
	"mime/multipart"
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"strconv"
	"time"
)

func AddHouse(ownerID string, house *request.HousePost) (string, error) {
	a := &models.Address{}
	//find address from commune code
	err := a.FindAddress(house.CommuneCode)
	if err != nil {
		log.Println(err)
		return "", err
	}
	h := &models.House{
		OwnerID:        ownerID,
		HouseType:      house.HouseType,
		Price: 			house.Price,
		Unit:   		house.Unit,
		Address:        *a,
		CommuneCode:    house.CommuneCode,
		Infrastructure: house.Infrastructure,
		NearBy:         house.NearBy,
		WithOwner:      house.WithOwner,
		ImageLink:      nil,
		Header:         house.Header,
		View:           0,
		Like:           0,
		Rented:         false,
		Content:        house.Content,
		PostTime:       time.Now().Unix(),
		Activate:       false,
		ExpiredTime:    0,
	}
	return h.PutItem()
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

func GetHouse(id string) (response.House, error) {
	o := &models.House{}
	res, err := o.GetResponse(id)
	if err != nil {
		return response.House{}, err
	}
	return res, nil
}

func FilterSearchResult(res []response.House, provinceID, districtID, communeID string) ([]response.House, error) {
	list := []response.House{}
	if communeID != "" {
		commune := models.Commune{}
		err := commune.GetItem(communeID)
		if err != nil {
			return []response.House{}, err
		}
		for _, i := range res {
			if i.Address.Commune == commune.Name {
				list = append(list, i)
			}
		}
	} else if districtID != "" {
		district := models.District{}
		err := district.GetItem(districtID)
		if err != nil {
			return []response.House{}, err
		}
		for _, i := range res {
			if i.Address.District == district.Name {
				list = append(list, i)
			}
		}
	} else if provinceID != "" {
		province := models.Province{}
		err := province.GetItem(provinceID)
		if err != nil {
			return []response.House{}, err
		}
		for _, i := range res {
			if i.Address.Province == province.Name {
				list = append(list, i)
			}
		}
	} else {
		return res, nil
	}
	return list, nil
}

func ViewHouse(id string) (error) {
	o := &models.House{}
	err := o.GetFromKey(id)
	if err != nil {
		return response.NotExisted
	}

	//increase view
	o.View++
	last := time.Unix(o.LastViewed, 0)
	if last.Month() == time.Now().Month() {
		o.MonthlyView++
	} else {
		o.MonthlyView = 1;
	}

	//create new one day//hour//views
	h, _, _ := time.Now().Clock()
	//create address
	c := &models.Commune{}
	err = c.GetItem(o.CommuneCode)
	if err != nil {
		return err
	}
	d := &models.District{}
	err = d.GetItem(c.ParentCode)
	if err != nil {
		return err
	}
	//create price range
	pm := o.Price
	switch o.Unit {
	case models.Month:
	case models.Quarter:
		pm = pm/3
	case models.Year:
		pm = pm/12
	}
	pr := models.PriceRangeFactory(pm)
	//update statistic
	stat := &models.Statistic{}
	err = stat.GetFromKey(stat.GetKeyNow())
	//check if statistic existed
	if err != nil {
		//create new statistic
		stat = &models.Statistic{
			ViewTime: map[string]map[string]int64{
				strconv.Itoa(time.Now().Day()) : {
					strconv.Itoa(h): 1,
				},
			},
			ViewLocation: map[string]map[string]int64{
				d.ParentCode : {
					d.Code : 1,
				},
			},
			ViewPriceRange: map[string]int64{
				string(pr) : 1,
			},
		}
	} else {
		//increase view
		//check if view time existed
		if v, ok := stat.ViewTime[strconv.Itoa(time.Now().Day())]; ok {
			//exist day
			if k, o := v[strconv.Itoa(h)]; o {
				//exist hour
				v[strconv.Itoa(h)] = k + 1
			}
			v[strconv.Itoa(h)] = 1
		} else {
			//not exist day
			stat.ViewTime[strconv.Itoa(time.Now().Day())] = map[string]int64{
				strconv.Itoa(h) : 1,
			}
		}
		//check if view location existed
		if v, ok := stat.ViewLocation[d.ParentCode]; ok {
			//exist province
			if k, o := v[d.Code]; o {
				//exist district
				v[d.Code] = k + 1
			}
			v[d.Code] = 1
		} else {
			//not exist province
			stat.ViewLocation[d.ParentCode] = map[string]int64{
				d.Code : 1,
			}
		}
		//check if view price existed
		if v, ok := stat.ViewPriceRange[string(pr)]; ok {
			//exist price range
			stat.ViewPriceRange[string(pr)] = v + 1
		} else {
			stat.ViewPriceRange[string(pr)] = 1
		}
	}
	//update statistic
	err = stat.PutItem()
	if err != nil {
		return err
	}
	_, err = o.PutItem()
	return err
}

func GetAllWaitHouse() ([]string, error) {
	h := &models.House{}
	list, err := h.GetAllWaitList()
	if err != nil {
		return []string{}, err
	}
	return list, nil
}

func GetPageWaitHouse(page int, count int) ([]string, error) {
	h := &models.House{}
	list, err := h.GetPaginateWaitList(page, count)
	if err != nil {
		return []string{}, err
	}
	return list, nil
}

func GetAllHouse() ([]response.House, error) {
	o := &models.House{}
	list, err := o.GetAllActivate()
	if err != nil {
		return []response.House{}, err
	}
	return list, nil
}

func GetPageHouse(page, count int) ([]response.House, error) {
	o := &models.House{}
	list, err :=  o.GetPageActivate(page, count)
	if err != nil {
		return []response.House{}, err
	}
	return list, nil
}

func GetAllHouseOfOwner(userID string) ([]response.House, error) {
	o := &models.House{}
	list, err := o.GetAllHouseOfOwner(userID)
	if err != nil {
		return []response.House{}, err
	}
	return list, nil
}

func GetPageHouseOfOwner(ownerID string, page int, count int) ([]response.House, error) {
	o := &models.House{}
	list, err := o.GetPaginateHouseOfUser(ownerID, page, count)
	if err != nil {
		return []response.House{}, err
	}
	return list, nil
}

func UpdateHouse(id string, ob *request.HousePut) error {
	h := &models.House{}
	err := h.GetFromKey(id)
	if err != nil {
		return err
	}
	a := &models.Address{}
	err = a.FindAddress(ob.CommuneCode)
	if err != nil {
		return err
	}
	h.CommuneCode = ob.CommuneCode
	h.Content = ob.Content
	h.Header = ob.Header
	h.WithOwner = ob.WithOwner
	h.NearBy = ob.NearBy
	h.Infrastructure = ob.Infrastructure
	h.Address = *a
	h.Price = ob.Price
	h.Unit = ob.Unit
	h.HouseType = ob.HouseType
	return h.UpdateItem(id)
}

func SearchHouse(key, provinceID, districtID, communeID string) ([]response.House, error) {
	h := &models.House{}
	res, err := h.SearchAllItem(key)
	if err != nil {
		return []response.House{}, err
	}
	return	FilterSearchResult(res, provinceID, districtID, communeID)
}

func SearchPageHouse(key, provinceID, districtID, communeID string, page, count int) ([]response.House, error) {
	h := &models.House{}
	res, err := h.SearchPaginateItem(key, page, count)
	if err != nil {
		return []response.House{}, err
	}
	return	FilterSearchResult(res, provinceID, districtID, communeID)
}

func DeleteHouse(id string) error {
	u := &models.House{}
	err := u.GetFromKey(id)
	if err != nil {
		return err
	}
	return u.Delete(id)
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
	if err != nil {
		return err
	}
	return nil
}

func RemoveFromFavourite(renterID, houseID string) error {
	r := &models.Renter{}
	err := r.GetFromKey(renterID)
	if err != nil {
		return err
	}
	list := []string{}
	//check if exist
	for _, i := range r.ListFavourite {
		if i == houseID {
			continue
		}
		list = append(list, i)
	}
	//add to list
	r.ListFavourite = list
	err = r.PutItem()
	if err != nil {
		return err
	}
	//decrease the number of likes
	h := &models.House{}
	err = h.GetFromKey(houseID)
	if err != nil {
		return err
	}
	h.Like--
	_, err = h.PutItem()
	return nil
}

