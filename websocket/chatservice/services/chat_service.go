package services

import (
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice/models"
	"time"
)

func GetMessageOfOwner(ownerID string, page, count int) (models.ResMessageConversation, int, error) {
	mes := models.MessageConversation{}
	if page < 0 || count < 0 {
		return models.ResMessageConversation{}, 0, response.BadRequest
	}
	if count != 0 {
		return mes.GetPaginateRecentOfOwner(ownerID, page, count)
	}
	return mes.GetAllByTimeOfOwner(ownerID)
}

func GetChattingOwner(page, length int) ([]response.Owner, int, error) {
	if page < 0 || length < 0 {
		return nil, 0, response.BadRequest
	}
	mes := models.MessageConversation{}
	if length != 0 {
		return mes.GetChattingOwner(page, length)
	}
	return mes.GetAllChattingOwner()
}

func AdminSendMessage(adminID string, request models.AdminMessage) (error) {
	bc := &models.BroadCastToOwner{
		AdminID:      adminID,
		SendTime:     time.Now().Unix(),
		Type:         models.ADMIN_MESSAGE,
		AdminMessage: request,
	}

	err := bc.PutItem()
	if err != nil {
		return err
	}
	if models.BcOwner[request.OwnerID] != nil {
		models.BcOwner[request.OwnerID] <- *bc
	}
	return nil
}