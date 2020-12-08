package response

type Owner struct {
	OwnerName   string  `json:"owner_name"`
	OwnerFullName string  `json:"owner_full_name"`
	Profile   Profile `json:"profile"`
	Address   Address `json:"address"`
	AverageStar float32 `json:"average_star"`
	Activate  bool	  `json:"activate"`
}

type Profile struct {
	IDCard      string `json:"id_card"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
