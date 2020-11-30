package request

import "rent-house/models"

type HousePost struct {
	HouseType      	   models.HouseType      `json:"house_type"`
	Price		   	   int     	      		 `json:"price"`
	Unit 		   	   models.Unit			 `json:"unit"`
	CommuneCode        string         		 `json:"commune_code"`
	Infrastructure 	   models.Infrastructure `json:"infrastructure"`
	NearBy         	   []string       		 `json:"near_by"`
	WithOwner      	   bool           		 `json:"with_owner"`
	Header         	   string         		 `json:"header"`
	Content        	   string        	     `json:"content"`
}

type HousePut struct {
	HouseType      models.HouseType      `json:"house_type"`
	Price		   int     	      		 `json:"price"`
	Unit 		   models.Unit			 `json:"unit"`
	CommuneCode    string         		 `json:"commune_code"`
	Infrastructure models.Infrastructure `json:"infrastructure"`
	NearBy         []string       		 `json:"near_by"`
	WithOwner      bool           		 `json:"with_owner"`
	Header         string         		 `json:"header"`
	Content        string         		 `json:"content"`
}

