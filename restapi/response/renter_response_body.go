package response

type Renter struct {
	RenterID 	    string		`json:"renter_id"`
	RenterName 		string 		`json:"renter_name"`
	PhoneNumber		string 		`json:"phone_number"`
	Email			string 		`json:"email"`
	ListFavourite	[]string	`json:"list_favourite"`
}

type RenterInfo struct {
	RenterID 	    string		`json:"renter_id"`
	RenterName 		string 		`json:"renter_name"`
}