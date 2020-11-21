package response

type House struct {
	HouseID		   string		  `json:"house_id"`
	OwnerID        string         `json:"owner_id"`
	HouseType      HouseType      `json:"house_type"`
	PricePerMonth  int            `json:"price"`
	PricePerYear   int            `json:"price_per_year"`
	Address        Address        `json:"address"`
	Infrastructure Infrastructure `json:"infrastructure"`
	NearBy         []string       `json:"near_by"`
	WithOwner      bool           `json:"with_owner"`
	ImageLink      []string       `json:"image_link"`
	Header         string         `json:"header"`
	View		   int 			  `json:"view"`
	Like		   int			  `json:"like"`
	Rented		   bool			  `json:"rented"`
	Content        string         `json:"content"`
	PostTime	   int64  		  `json:"post_time"`
	Activate	   bool  		  `json:"activate"`
	ExpiredTime	   int64  		  `json:"expired_time"`
}

type HouseSearch struct {
	ObjectID string `json:"objectID"`
	House
}

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
