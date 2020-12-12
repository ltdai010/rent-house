package models

type OwnerMessage struct {
	Message  string `json:"message"`
}

type AdminMessage struct {
	OwnerID	 string	`json:"owner_id"`
	Message  string `json:"message"`
}