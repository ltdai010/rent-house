package request

type RenterPost struct {
	RenterName 		string 			`json:"renter_name"`
	RenterFullName 		string 		`json:"renter_full_name"`
	Password		string			`json:"password"`
	PhoneNumber		string 			`json:"phone_number"`
	Email			string 			`json:"email"`
}

type RenterPut struct {
	RenterFullName 		string 		`json:"renter_full_name"`
	Password		string			`json:"password"`
	PhoneNumber		string 			`json:"phone_number"`
	Email			string 			`json:"email"`
}
