package statisticservices

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func MostViewInMonth() (house response.House, err error) {
	h := &models.House{}
	res, err := h.GetMaxViewHouseInMonth()
	return res, err
}

func ViewInHourThisMonth() (mapTime map[string]int64, err error) {
	stat := &models.Statistic{}
	mapTime = map[string]int64{}
	err = stat.GetFromKey(stat.GetKeyNow())
	if err != nil {
		return nil, err
	}
	for _, dayView := range stat.ViewTime {
		for i, j := range dayView {
			mapTime[i] += j
		}
	}
	return
}

func TimelineViewThisMonth() (mapTime map[string]int64, err error) {
	stat := &models.Statistic{}
	mapTime = map[string]int64{}
	err = stat.GetFromKey(stat.GetKeyNow())
	if err != nil {
		return nil, err
	}
	for d, dayView := range stat.ViewTime {
		for _, j := range dayView {
			mapTime[d] += j
		}
	}
	return
}
