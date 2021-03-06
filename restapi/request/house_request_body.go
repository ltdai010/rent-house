package request

import "rent-house/models"

type HousePost struct {
	HouseType      	   models.HouseType      `json:"house_type"`
	Price		   	   float64	      		 `json:"price"`
	Unit 		   	   models.Unit			 `json:"unit"`
	PreOrder	   	   int	  	     		 `json:"pre_order"`
	Surface		   	   int			         `json:"surface"`
	CommuneCode        string         		 `json:"commune_code"`
	Street  		   string				 `json:"street"`
	Infrastructure 	   models.Infrastructure `json:"infrastructure"`
	ImageLink          []string       		 `json:"image_link"`
	NearBy         	   []string       		 `json:"near_by"`
	WithOwner      	   bool           		 `json:"with_owner"`
	AppearTime	       int64		  		 `json:"appear_time"`
	Header         	   string         		 `json:"header"`
	Content        	   string        	     `json:"content"`
}

type HousePut struct {
	HouseType      models.HouseType      `json:"house_type"`
	Price		   float64 	      		 `json:"price"`
	Unit 		   models.Unit			 `json:"unit"`
	CommuneCode    string         		 `json:"commune_code"`
	Street  	   string				 `json:"street"`
	PreOrder	   int	  	     		 `json:"pre_order"`
	Surface		   int			         `json:"surface"`
	Infrastructure models.Infrastructure `json:"infrastructure"`
	ImageLink      []string       		 `json:"image_link"`
	NearBy         []string        		 `json:"near_by"`
	WithOwner      bool           		 `json:"with_owner"`
	Header         string         		 `json:"header"`
	Content        string         		 `json:"content"`
}

type DeniedComment struct {
	HouseID string `json:"house_id"`
	Comment string `json:"comment"`
}

