package services

import (
	"rent-house/restapi/response"
	"rent-house/websocket/chatservice/models"
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
