package request

import "rent-house/models"

type PostOwner struct {
	OwnerName string  		 `json:"owner_name"`
	Profile   models.Profile `json:"profile"`
	Address   models.Address `json:"address"`
	Activate  bool	  		 `json:"activate"`
}

type PutOwner struct {
	OwnerName string  		 `json:"owner_name"`
	Profile   models.Profile `json:"profile"`
	Address   models.Address `json:"address"`
	Activate  bool	  		 `json:"activate"`
}