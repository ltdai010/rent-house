package models

import (
	"cloud.google.com/go/firestore"
	"log"
	"rent-house/models"
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice"
)

type MessageConversation struct {
	MapMessage map[string]interface{}	`json:"map_message"`
}

type ResMessageConversation struct {
	Messages []interface{}	`json:"messages"`
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
	list, err := this.GetCollection().Documents(models.Ctx).GetAll()
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
	res.Messages = []interface{}{}
	doc, err := this.GetCollection().Doc(ownerID).Get(models.Ctx)
	if err != nil {
		return ResMessageConversation{}, 0, err
	}
	err = doc.DataTo(mc)
	if err != nil {
		return ResMessageConversation{}, 0, err
	}
	if mc.MapMessage != nil {
		for _, i := range mc.MapMessage {
			res.Messages = append(res.Messages, i)
		}
	}
	return res, len(res.Messages), nil
}

func (this *MessageConversation) GetPaginateRecentOfOwner(ownerID string, page, count int) (ResMessageConversation, int, error) {
	mc := MessageConversation{}
	res := ResMessageConversation{}
	res.Messages = []interface{}{}
	start := page * count
	end := start + count
	tmp := start
	doc, err := this.GetCollection().Doc(ownerID).Get(models.Ctx)
	if err != nil {
		log.Println(err, " websocket/chatservice/models/model.go:95")
		return ResMessageConversation{}, 0, err
	}
	err = doc.DataTo(&mc)
	if err != nil {
		log.Println(err, " websocket/chatservice/models/model.go:100")
		return ResMessageConversation{}, 0, err
	}
	list := []interface{}{}
	if mc.MapMessage != nil {
		for _, i := range mc.MapMessage {
			list = append(list, i)
		}
	}
	total := len(list)
	log.Println(total)
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
