package models

import (
	"cloud.google.com/go/firestore"
	"github.com/gorilla/websocket"
	"net/http"
	"rent-house/models"
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice"
)

var (
	//connected client
	Clients  = make(map[string]*websocket.Conn)     // connected Clients
	//admin receiver
	Admin    = make(map[string]*websocket.Conn)
	//broadcast to admin channel
	BcAdmin = make(map[string]chan BroadCastToAdmin) // broadcastbody channel
	//broadcast to client channel
	BcOwner = make(map[string]chan BroadCastToOwner)
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type MessageConversation struct {
	Messages []Message	`json:"map_message"`
	LatestMsgTime int64	`json:"latest_msg_time"`
}

type ResMessageConversation struct {
	Messages []Message	`json:"messages"`
}

type Message struct {
	AdminID  string `json:"admin_id"`
	SendTime int64  `json:"send_time"`
	Type 	MessageType `json:"type"`
	OwnerID	 string	`json:"owner_id"`
	Message  string `json:"message"`
	ImageLink string `json:"image_link"`
}

type MessageType string

const (
	OWNER_MESSAGE = "owner_message"
	ADMIN_MESSAGE = "admin_message"
)

func (this *MessageConversation) GetCollectionKey() string {
	return chatservice.CHAT
}

func (this *MessageConversation) GetCollection() *firestore.CollectionRef {
	return models.Client.Collection(this.GetCollectionKey())
}

func (this *MessageConversation) GetAllChattingOwner() ([]response.Owner, int, error) {
	res := []response.Owner{}
	list, err := this.GetCollection().OrderBy("LatestMsgTime", firestore.Desc).Documents(models.Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(list)
	for _, i := range list {
		owner := models.Owner{}
		err = owner.GetFromKey(i.Ref.ID)
		if err != nil {
			continue
		}
		res = append(res, response.Owner{
			OwnerName:     owner.OwnerName,
			OwnerFullName: owner.OwnerFullName,
			Profile:       response.Profile{
				IDCard:      owner.Profile.IDCard,
				PhoneNumber: owner.Profile.PhoneNumber,
				Email:       owner.Profile.Email,
			},
			Address:       response.Address{
				Province: owner.Address.Province,
				District: owner.Address.District,
				Commune:  owner.Address.Commune,
				Street:   owner.Address.Street,
			},
			PostTime: owner.PostTime,
			Activate: owner.Activate,
		})
	}
	return res, total, nil
}

func (this *MessageConversation) GetChattingOwner(page, length int) ([]response.Owner, int, error) {
	res := []response.Owner{}
	start := page * length
	end := start + length
	list, err := this.GetCollection().Documents(models.Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(list)
	if start > total {
		return nil, 0, response.BadRequest
	}
	if end > total {
		end = total
	}
	for _, i := range list[start : end] {
		owner := models.Owner{}
		err = owner.GetFromKey(i.Ref.ID)
		if err != nil {
			continue
		}
		res = append(res, response.Owner{
			OwnerName:     owner.OwnerName,
			OwnerFullName: owner.OwnerFullName,
			Profile:       response.Profile{
				IDCard:      owner.Profile.IDCard,
				PhoneNumber: owner.Profile.PhoneNumber,
				Email:       owner.Profile.Email,
			},
			Address:       response.Address{
				Province: owner.Address.Province,
				District: owner.Address.District,
				Commune:  owner.Address.Commune,
				Street:   owner.Address.Street,
			},
			PostTime: owner.PostTime,
			Activate: owner.Activate,
		})
	}
	return res, total, nil
}

func (this *MessageConversation) GetAllByTimeOfOwner(ownerID string) (ResMessageConversation, int, error) {
	mc := MessageConversation{}
	res := ResMessageConversation{}
	res.Messages = []Message{}
	doc, err := this.GetCollection().Doc(ownerID).Get(models.Ctx)
	if err != nil {
		return ResMessageConversation{}, 0, err
	}
	err = doc.DataTo(&mc)
	if err != nil {
		return ResMessageConversation{}, 0, err
	}
	if mc.Messages != nil {
		for _, i := range mc.Messages {
			res.Messages = append(res.Messages, i)
		}
	}
	return res, len(res.Messages), nil
}

func (this *MessageConversation) GetPaginateRecentOfOwner(ownerID string, page, count int) (ResMessageConversation, int, error) {
	mc := MessageConversation{}
	res := ResMessageConversation{}
	res.Messages = []Message{}
	start := page * count
	end := start + count
	tmp := start
	doc, err := this.GetCollection().Doc(ownerID).Get(models.Ctx)
	if err != nil {
		return ResMessageConversation{}, 0, err
	}
	err = doc.DataTo(&mc)
	if err != nil {
		return ResMessageConversation{}, 0, err
	}
	list := []Message{}
	if mc.Messages != nil {
		for _, i := range mc.Messages {
			list = append(list, i)
		}
	}
	total := len(list)
	start = total - end
	end = total - tmp
	if end < 0 {
		return ResMessageConversation{}, 0, response.BadRequest
	}
	if start < 0 {
		start = 0
	}
	res.Messages = list[start : end]
	return res, total, nil
}
