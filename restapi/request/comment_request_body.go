package request

type CommentPost struct {
	Content   string `json:"content"`
	Header    string `json:"header"`
}

type CommentPut struct {
	Content   string `json:"content"`
	Header    string `json:"header"`
}
