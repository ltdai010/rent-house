package models

import (
	"cloud.google.com/go/firestore"
	"rent-house/models"
	"rent-house/websocket/chatservice"
)

type BroadCastToOwner struct {
	AdminID  string `json:"admin_id"`
	SendTime int64  `json:"send_time"`
	Type 	MessageType `json:"type"`
	AdminMessage
}


func (this *BroadCastToOwner) GetCollectionKey() string {
	return chatservice.CHAT
}

func (this *BroadCastToOwner) GetCollection() *firestore.CollectionRef {
	return models.Client.Collection(this.GetCollectionKey())
}

func (this *BroadCastToOwner) PutItem() error {
	mc := &MessageConversation{
		Messages: []Message{},
	}
	doc, err := this.GetCollection().Doc(this.OwnerID).Get(models.Ctx)
	if err != nil {
		mc.Messages = append(mc.Messages, Message{
			AdminID:   this.AdminID,
			SendTime:  this.SendTime,
			Type:      this.Type,
			OwnerID:   this.OwnerID,
			Message:   this.Message,
			ImageLink: this.ImageLink,
		})
	} else {
		err = doc.DataTo(mc)
		if err != nil {
			return err
		}
		mc.Messages = append(mc.Messages, Message{
			AdminID:   this.AdminID,
			SendTime:  this.SendTime,
			Type:      this.Type,
			OwnerID:   this.OwnerID,
			Message:   this.Message,
			ImageLink: this.ImageLink,
		})
		mc.LatestMsgTime = this.SendTime
	}
	_, err = this.GetCollection().Doc(this.OwnerID).Set(models.Ctx, mc)
	if err != nil {
		return err
	}
	return err
}

