package models

import (
	"google.golang.org/api/iterator"
	"rent-house/consts"
)

type Status string

const (
	InActivated = "inactivated"
	Activated = "activated"
	Denied = "denied"
	Extend = "extend"
)

type Address struct {
	Province  string `json:"province"`
	District  string `json:"district"`
	Commune   string `json:"commune"`
	Street    string `json:"street"`
}

func (this *Province) GetAll() ([]Province,error) {
	iter := Client.Collection(consts.PROVINCE).Documents(Ctx)
	list := []Province{}
	for {
		p := Province{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []Province{}, err
		}
		err = doc.DataTo(&p)
		if err != nil {
			continue
		}
		list = append(list, p)
	}
	return list, nil
}

func (this *District) GetAll(provinceID string) ([]District,error) {
	iter := Client.Collection(consts.DISTRICT).Where("ParentCode", "==", provinceID).Documents(Ctx)
	list := []District{}
	for {
		p := District{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []District{}, err
		}
		err = doc.DataTo(&p)
		if err != nil {
			continue
		}
		list = append(list, p)
	}
	return list, nil
}

func (this *Commune) GetAll(districtID string) ([]Commune,error) {
	iter := Client.Collection(consts.COMMUNE).Where("ParentCode", "==", districtID).Documents(Ctx)
	list := []Commune{}
	for {
		p := Commune{}
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return []Commune{}, err
		}
		err = doc.DataTo(&p)
		if err != nil {
			continue
		}
		list = append(list, p)
	}
	return list, nil
}

func (this *Province) GetItem(provinceID string) error {
	doc, err := Client.Collection(consts.PROVINCE).Doc(provinceID).Get(Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *District) GetItem(districtID string) error {
	doc, err := Client.Collection(consts.DISTRICT).Doc(districtID).Get(Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *Commune) GetItem(communeID string) error {
	doc, err := Client.Collection(consts.COMMUNE).Doc(communeID).Get(Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *Address) FindAddress(communeCode string) (error) {
	//get the commune
	doc, err := Client.Collection(consts.COMMUNE).Doc(communeCode).Get(Ctx)
	if err != nil {
		return err
	}
	c := &Commune{}
	err = doc.DataTo(c)
	if err != nil {
		return err
	}
	//get the commune's district
	doc, err = Client.Collection(consts.DISTRICT).Doc(c.ParentCode).Get(Ctx)
	if err != nil {
		return err
	}
	d := &District{}
	err = doc.DataTo(d)
	if err != nil {
		return err
	}
	//get the district's province
	doc, err = Client.Collection(consts.PROVINCE).Doc(d.ParentCode).Get(Ctx)
	if err != nil {
		return err
	}
	p := &Province{}
	err = doc.DataTo(p)
	if err != nil {
		return err
	}
	this.Province = p.Name
	this.District = d.Name
	this.Commune = c.Name
	return  err
}

type Infrastructure struct {
	PrivateBathroom bool   `json:"private_bathroom"`
	Heater          bool   `json:"heater"`
	AirCondition    bool   `json:"air_condition"`
	Balcony         bool   `json:"balcony"`
	ElectricPrice   int    `json:"electric_price"`
	WaterPrice      int    `json:"water_price"`
	NumberOfRoom	int	   `json:"number_of_room"`
	Kitchen         bool   `json:"kitchen"`
	Other           string `json:"other"`
}

type Unit int

const (
	Month = iota
	Quarter
	Year
)

type HouseType int

const (
	Room = iota
	MiniApartment
	FullHouse
	Apartment
)

type PriceRange string

const (
	VeryLow = "* - 500"
	DownLow = "500 - 700"
	UpLow     = "700 - 1000"
	DownMedium = "1000 - 1500"
	UpMedium = "1500 - 2000"
	High = "2000 - 2500"
	VeryHigh = "2500 - 3500"
	ExHigh = "3500 - 5000"
	More = "5000 - *"
)

func (p PriceRange) ToRange() (int, int) {
	switch p {
	case VeryLow:
		return 0, 500
	case DownLow:
		return 500, 700
	case UpLow:
		return 700, 1000
	case DownMedium:
		return 1000, 1500
	case UpMedium:
		return 1500, 2000
	case High:
		return 2000, 2500
	case VeryHigh:
		return 2500, 3500
	case ExHigh:
		return 3500, 5000
	default:
		return 5000, 99999999
	}
}

func PriceRangeFactory(price float64) PriceRange {
	if price < 500 {
		return VeryLow
	} else if price >= 500 && price < 700 {
		return DownLow
	} else if price >= 700 && price < 1000 {
		return UpLow
	} else if price >= 1000 && price < 1500 {
		return DownMedium
	} else if price >= 1500 && price < 2000 {
		return UpMedium
	} else if price >= 2000 && price < 2500 {
		return High
	} else if price >= 2500 && price < 3500 {
		return VeryHigh
	} else if price >= 3500 && price < 5000 {
		return ExHigh
	}
	return More
}

type PostTime struct {
	HouseID		string
	PostTime	int64
	ExpireTime	int64
}

type Profile struct {
	IDCard      string `json:"id_card"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type Login struct {
	Username 		string  	`json:"username"`
	Password		string		`json:"password"`
}

type Province struct {
	Name	string  		`json:"name"`
	Slug	string			`json:"slug"`
	Type	string			`json:"type"`
	NameWithType	string  `json:"name_with_type"`
	Code 			string	`json:"code"`
}

type District struct {
	Name			string	`json:"name"`
	Type			string	`json:"type"`
	Slug 			string	`json:"slug"`
	NameWithType	string	`json:"name_with_type"`
	Path			string	`json:"path"`
	PathWithType	string	`json:"path_with_type"`
	Code 			string	`json:"code"`
	ParentCode		string	`json:"parent_code"`
}

type Commune struct {
	Name			string	`json:"name"`
	Type			string	`json:"type"`
	Slug			string	`json:"slug"`
	NameWithType	string	`json:"name_with_type"`
	Path 			string	`json:"path"`
	PathWithType	string	`json:"path_with_type"`
	Code 			string	`json:"code"`
	ParentCode		string	`json:"parent_code"`
}