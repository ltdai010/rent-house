package models

import (
	"google.golang.org/api/iterator"
	"rent-house/consts"
)

type Address struct {
	Province  string `json:"province"`
	District  string `json:"district"`
	Commune   string `json:"commune"`
}

func (this *Province) GetAll() ([]Province,error) {
	iter := client.Collection(consts.PROVINCE).Documents(ctx)
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
	iter := client.Collection(consts.DISTRICT).Where("ParentCode", "==", provinceID).Documents(ctx)
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
	iter := client.Collection(consts.COMMUNE).Where("ParentCode", "==", districtID).Documents(ctx)
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
	doc, err := client.Collection(consts.PROVINCE).Doc(provinceID).Get(ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *District) GetItem(districtID string) error {
	doc, err := client.Collection(consts.DISTRICT).Doc(districtID).Get(ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *Commune) GetItem(communeID string) error {
	doc, err := client.Collection(consts.COMMUNE).Doc(communeID).Get(ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *Address) FindAddress(communeCode string) (error) {
	//get the commune
	doc, err := client.Collection(consts.COMMUNE).Doc(communeCode).Get(ctx)
	if err != nil {
		return err
	}
	c := &Commune{}
	err = doc.DataTo(c)
	if err != nil {
		return err
	}
	//get the commune's district
	doc, err = client.Collection(consts.DISTRICT).Doc(c.ParentCode).Get(ctx)
	if err != nil {
		return err
	}
	d := &District{}
	err = doc.DataTo(d)
	if err != nil {
		return err
	}
	//get the district's province
	doc, err = client.Collection(consts.PROVINCE).Doc(d.ParentCode).Get(ctx)
	if err != nil {
		return err
	}
	p := &Province{}
	err = doc.DataTo(p)
	if err != nil {
		return err
	}
	this.Province = p.NameWithType
	this.District = d.NameWithType
	this.Commune = c.NameWithType
	return  err
}

type Infrastructure struct {
	PrivateBathroom bool   `json:"private_bathroom"`
	Heater          bool   `json:"heater"`
	AirCondition    bool   `json:"air_condition"`
	Balcony         bool   `json:"balcony"`
	ElectricPrice   int    `json:"electric_price"`
	WaterPrice      int    `json:"water_price"`
	Other           string `json:"other"`
}

type HouseType int

const (
	Room = iota
	MiniApartment
	FullHouse
	Apartment
)

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