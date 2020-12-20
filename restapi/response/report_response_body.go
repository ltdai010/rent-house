package response

type Report struct {
	ReportID string	`json:"report_id"`
	Tittle   string `json:"tittle"`
	Content  string `json:"content"`
	RenterID string `json:"renter_name"`
	House    House  `json:"house"`
	Seen	 bool	`json:"seen"`
	SendTime int64  `json:"send_time"`
}
