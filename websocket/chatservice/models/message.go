package models

type Message struct {
	OwnerID			string `json:"owner_id"`
	SendTime		int64  `json:"send_time"`
	MessageContent	string `json:"message_content"`
}