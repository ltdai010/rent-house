package models


type Address struct {
	Province string `json:"province"`
	District string `json:"district"`
	Street   string `json:"street"`
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