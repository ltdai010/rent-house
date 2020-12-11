package models

type Notification struct {
	Content    string `json:"content"`
	ReceiverID string `json:"receiver_id"`
	SendTime   int64  `json:"send_time"`
}
