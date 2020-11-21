package request

import "rent-house/models"

type OwnerPost struct {
	OwnerName 		string  	`json:"owner_name"`
	Password		string		`json:"password"`
	OwnerFullName	string		`json:"owner_full_name"`
	Profile   		models.Profile 	`json:"profile"`
	Address   		models.Address 	`json:"address"`
}

type OwnerPut struct {
	Password		string		`json:"password"`
	OwnerFullName	string		`json:"owner_full_name"`
	Profile   		models.Profile 	`json:"profile"`
	Address   		models.Address 	`json:"address"`
}
