package models

import "rent-house/consts"

type Address struct {
	Province  string `json:"province"`
	District  string `json:"district"`
	Commune   string `json:"commune"`
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