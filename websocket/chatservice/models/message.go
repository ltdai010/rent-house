package models

type BroadCastToAdmin struct {
	OwnerID			string `json:"owner_id"`
	SendTime		int64  `json:"send_time"`
	OwnerMessage
}

type BroadCastToOwner struct {
	AdminID  string `json:"admin_id"`
	SendTime int64  `json:"send_time"`
	AdminMessage
}
