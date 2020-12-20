package response

type Comment struct {
	CommentID string `json:"comment_id"`
	Content   string `json:"content"`
	RenterName  string `json:"renter_name"`
	HouseID   string `json:"house_id"`
	PostTime  int64  `json:"post_time"`
	Star      int    `json:"star"`
	Activate  bool   `json:"activate"`
}
