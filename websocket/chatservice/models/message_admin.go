package models

import (
	"cloud.google.com/go/firestore"
	"rent-house/models"
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice"
)

type BroadCastToAdmin struct {
	OwnerID			string `json:"owner_id"`
	SendTime		int64  `json:"send_time"`
	OwnerMessage
}

type ResBroadCastToAdmin struct {
	ID				string `json:"id"`
	OwnerID			string `json:"owner_id"`
	SendTime		int64  `json:"send_time"`
	OwnerMessage
}


func (this *BroadCastToAdmin) GetCollectionKey() string {
	return chatservice.ADMIN_CHAT_RECEIVER
}

func (this *BroadCastToAdmin) GetCollection() *firestore.CollectionRef {
	return models.Client.Collection(this.GetCollectionKey())
}

func (this *BroadCastToAdmin) PutItem() error {
	_, _, err := this.GetCollection().Add(models.Ctx, this)
	return err
}

func (this *BroadCastToAdmin) GetFromKey(key string) error {
	doc, err := this.GetCollection().Doc(key).Get(models.Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *BroadCastToAdmin) UpdateItem(key string) error {
	_, err := this.GetCollection().Doc(key).Set(models.Ctx, this)
	return err
}

func (this *BroadCastToAdmin) GetAllByTimeOfOwner(ownerID string) ([]ResBroadCastToAdmin, error) {
	res := []ResBroadCastToAdmin{}
	list, err := this.GetCollection().Where("OwnerID", "==", ownerID).OrderBy("SendTime", firestore.Desc).Documents(models.Ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range list {
		r :=  ResBroadCastToAdmin{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		r.ID = i.Ref.ID
		res = append(res, r)
	}
	return res, nil
}

func (this *BroadCastToAdmin) GetPaginateRecentOfOwner(ownerID string, page, count int) ([]ResBroadCastToAdmin, int, error) {
	res := []ResBroadCastToAdmin{}
	list, err := this.GetCollection().Where("OwnerID", "==", ownerID).OrderBy("SendTime", firestore.Desc).Documents(models.Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(list)
	end := (page + 1) * count
	if page*count > total {
		return nil, 0, response.BadRequest
	}
	if end > total {
		end = total
	}
	for _, i := range list[page*count : end] {
		r := ResBroadCastToAdmin{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		r.ID = i.Ref.ID
		res = append(res, r)
	}
	return res, total, nil
}

