package response

type Comment struct {
	CommentID string `json:"comment_id"`
	Content   string `json:"content"`
	RenterID  string `json:"renter_id"`
	HouseID   string `json:"house_id"`
	PostTime  int64  `json:"post_time"`
	Star      int    `json:"star"`
	Activate  bool   `json:"activate"`
}
