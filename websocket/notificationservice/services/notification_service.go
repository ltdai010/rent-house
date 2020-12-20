package services

import (
	"rent-house/restapi/response"
	"rent-house/websocket/notificationservice/models"
)

func GetAllNotificationOfOwner(ownerID string) ([]models.ResNotification, error) {
	notice := models.Notification{}
	list, err := notice.GetAllByTimeOfOwner(ownerID)
	go SeenMultiple(list)
	return list, err
}

func GetPageNotificationOfOwner(ownerID string, page, count int) ([]models.ResNotification, int, error) {
	if page < 0 || count < 0 {
		return nil, 0, response.BadRequest
	}
	notice := models.Notification{}
	list, total, err := notice.GetPaginateRecentOfOwner(ownerID, page, count)
	go SeenMultiple(list)
	return list, total, err
}

func SeenMultiple(list []models.ResNotification) {
	for _, i := range list {
		i.Seen = true
		not := &models.Notification{}
		err := not.GetFromKey(i.NotificationID)
		if err != nil {
			continue
		}
		go not.UpdateItem(i.NotificationID)
	}
}