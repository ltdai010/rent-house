package models

import (
	"cloud.google.com/go/firestore"
	"rent-house/models"
	"rent-house/restapi/response"
	"rent-house/websocket/notificationservice"
)

type Notification struct {
	Type     NotificationType `json:"type"`
	OwnerID  string           `json:"owner_id"`
	House    response.House   `json:"house_id"`
	SendTime int64            `json:"send_time"`
	Seen     bool             `json:"seen"`
}

type ResNotification struct {
	NotificationID string `json:"notification_id"`
	Notification
}

type NotificationType string

const (
	Denied = "denied"
	Activated = "activated"
)

func (this *Notification) GetCollectionKey() string {
	return notificationservice.NOTIFICATION
}

func (this *Notification) GetCollection() *firestore.CollectionRef {
	return models.Client.Collection(this.GetCollectionKey())
}

func (this *Notification) PutItem() error {
	_, _, err := this.GetCollection().Add(models.Ctx, this)
	return err
}

func (this *Notification) GetFromKey(key string) error {
	doc, err := this.GetCollection().Doc(key).Get(models.Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *Notification) UpdateItem(key string) error {
	_, err := this.GetCollection().Doc(key).Set(models.Ctx, this)
	return err
}

func (this *Notification) GetAllByTimeOfOwner(ownerID string) ([]ResNotification, error) {
	res := []ResNotification{}
	list, err := this.GetCollection().Where("OwnerID", "==", ownerID).OrderBy("SendTime", firestore.Desc).Documents(models.Ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range list {
		r := ResNotification{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		r.NotificationID = i.Ref.ID
		res = append(res, r)
	}
	return res, nil
}

func (this *Notification) GetPaginateRecentOfOwner(ownerID string, page, count int) ([]ResNotification, int, error) {
	res := []ResNotification{}
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
		r := ResNotification{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		r.NotificationID = i.Ref.ID
		res = append(res, r)
	}
	return res, total, nil
}
