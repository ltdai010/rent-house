package response

type Owner struct {
	OwnerID   string  `json:"owner_id"`
	OwnerName string  `json:"owner_name"`
	Profile   Profile `json:"profile"`
	Address   Address `json:"address"`
	Activate  bool	  `json:"activate"`
}

type Profile struct {
	IDCard      string `json:"id_card"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
