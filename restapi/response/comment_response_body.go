package response

type Comment struct {
	CommentID string `json:"comment_id"`
	Content   string `json:"content"`
	Header    string `json:"header"`
	PostTime  int64  `json:"post_time"`
	Activate  bool	 `json:"activate"`
}
