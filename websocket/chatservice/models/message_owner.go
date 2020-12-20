package models

import (
	"cloud.google.com/go/firestore"
	"rent-house/models"
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice"
)

type BroadCastToOwner struct {
	AdminID  string `json:"admin_id"`
	SendTime int64  `json:"send_time"`
	AdminMessage
}

type ResBroadCastToOwner struct {
	ID				string `json:"id"`
	AdminID  string `json:"admin_id"`
	SendTime int64  `json:"send_time"`
	AdminMessage
}


func (this *BroadCastToOwner) GetCollectionKey() string {
	return chatservice.OWNER_CHAT_RECEIVER
}

func (this *BroadCastToOwner) GetCollection() *firestore.CollectionRef {
	return models.Client.Collection(this.GetCollectionKey())
}

func (this *BroadCastToOwner) PutItem() error {
	_, _, err := this.GetCollection().Add(models.Ctx, this)
	return err
}

func (this *BroadCastToOwner) GetFromKey(key string) error {
	doc, err := this.GetCollection().Doc(key).Get(models.Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *BroadCastToOwner) UpdateItem(key string) error {
	_, err := this.GetCollection().Doc(key).Set(models.Ctx, this)
	return err
}

func (this *BroadCastToOwner) GetAllByTimeOfOwner(ownerID string) ([]ResBroadCastToOwner, error) {
	res := []ResBroadCastToOwner{}
	list, err := this.GetCollection().Where("OwnerID", "==", ownerID).OrderBy("SendTime", firestore.Desc).Documents(models.Ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range list {
		r :=  ResBroadCastToOwner{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		r.ID = i.Ref.ID
		res = append(res, r)
	}
	return res, nil
}

func (this *BroadCastToOwner) GetPaginateRecentOfOwner(ownerID string, page, count int) ([]ResBroadCastToOwner, int, error) {
	res := []ResBroadCastToOwner{}
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
		r := ResBroadCastToOwner{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		r.ID = i.Ref.ID
		res = append(res, r)
	}
	return res, total, nil
}
