package request

type CommentPost struct {
	Content  string `json:"content"`
	Header   string `json:"header"`
	Star     int    `json:"star"`
}

type CommentPut struct {
	Content   string `json:"content"`
	Header    string `json:"header"`
	Star      int    `json:"star"`
}
