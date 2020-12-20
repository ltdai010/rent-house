package models

type OwnerMessage struct {
	Message  string `json:"message"`
	ImageLink string `json:"image_link"`
}

type AdminMessage struct {
	OwnerID	 string	`json:"owner_id"`
	Message  string `json:"message"`
	ImageLink string `json:"image_link"`
}