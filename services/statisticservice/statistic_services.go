package statisticservice

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func MostViewInMonth() (house response.House, err error) {
	h := &models.House{}
	res, err := h.GetMaxViewHouseInMonth()
	return res, err
}

func ViewInHourThisMonth() (mapTime map[string]string, err error) {
	return
}

func TimelineViewThisMonth() (mapTime map[string]int, err error) {
	return
}
