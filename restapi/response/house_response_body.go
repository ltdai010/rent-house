package response

type House struct {
	HouseID		   string		  `json:"house_id"`
	OwnerID        string         `json:"owner_id"`
	HouseType      HouseType      `json:"house_type"`
	Price		   float64 	      `json:"price"`
	Unit 		   Unit			  `json:"unit"`
	Address        Address        `json:"address"`
	Infrastructure Infrastructure `json:"infrastructure"`
	NearBy         []string       `json:"near_by"`
	PreOrder	   int			  `json:"pre_order"`
	Surface		   int			  `json:"surface"`
	WithOwner      bool           `json:"with_owner"`
	ImageLink      []string       `json:"image_link"`
	LastViewed	   int64 		  `json:"last_viewed"`
	MonthlyView	   int			  `json:"monthly_view"`
	Header         string         `json:"header"`
	View		   int64		  `json:"view"`
	Like		   int64		  `json:"like"`
	Rented		   bool			  `json:"rented"`
	Content        string         `json:"content"`
	PostTime	   int64  		  `json:"post_time"`
	Activate	   bool  		  `json:"activate"`
	Review 		   map[string]int `json:"review"`
	AppearTime	   int64		  `json:"appear_time"`
	ExpiredTime	   int64  		  `json:"expired_time"`
}

type HouseSearch struct {
	ObjectID string `json:"objectID"`
	House
}

type Unit int

const (
	Month = iota
	Quarter
	Year
)

type Address struct {
	Province string `json:"province"`
	District string `json:"district"`
	Commune  string `json:"commune"`
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

type HouseType int

const (
	Room = iota
	MiniApartment
	FullHouse
	Apartment
)
