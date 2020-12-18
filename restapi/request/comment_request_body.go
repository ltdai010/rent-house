package request

type CommentPost struct {
	Content  string `json:"content"`
	Star     int    `json:"star"`
}

type CommentPut struct {
	Content   string `json:"content"`
	Star      int    `json:"star"`
}

