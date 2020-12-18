package request

type ReportPost struct {
	Tittle   string `json:"tittle"`
	Content  string `json:"content"`
	Seen	 bool	`json:"seen"`
	PostTime int64  `json:"post_time"`
}
