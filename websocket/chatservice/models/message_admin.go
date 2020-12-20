package models

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"rent-house/models"
	"rent-house/websocket/chatservice"
)

type BroadCastToAdmin struct {
	OwnerID			string `json:"owner_id"`
	SendTime		int64  `json:"send_time"`
	Type 			MessageType `json:"type"`
	OwnerMessage
}

func (this *BroadCastToAdmin) GetCollectionKey() string {
	return chatservice.CHAT
}

func (this *BroadCastToAdmin) GetCollection() *firestore.CollectionRef {
	return models.Client.Collection(this.GetCollectionKey())
}

func (this *BroadCastToAdmin) PutItem() error {
	mc := &MessageConversation{
		MapMessage: map[string]interface{}{},
	}
	doc, err := this.GetCollection().Doc(this.OwnerID).Get(models.Ctx)
	if err != nil {
		mc.MapMessage["0"] = *this
	} else {
		err = doc.DataTo(mc)
		if err != nil {
			return err
		}
		mc.MapMessage[fmt.Sprint(len(mc.MapMessage))] = *this
	}
	_, err = this.GetCollection().Doc(this.OwnerID).Set(models.Ctx, mc)
	if err != nil {
		return err
	}
	return err
}




