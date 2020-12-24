package statisticservices

import (
	"fmt"
	"rent-house/models"
	"rent-house/restapi/response"
	"time"
)

func MostViewInMonth(length int) (house response.House, err error) {
	h := &models.House{}
	res, err := h.GetMaxViewHouseInMonth(length)
	return res, err
}

func TimelineViewThisMonth() (map[string]map[string]int64, error) {
	stat := &models.Statistic{}
	mapTime := MakeMonthMapNow()
	err := stat.GetFromKey(stat.GetKeyNow())
	if err != nil {
		return mapTime, err
	}
	for d, dayView := range stat.ViewTime {
		for h, j := range dayView {
			mapTime[d][h] += j
		}
	}
	return mapTime, nil
}

func MakeMonthMapNow() map[string]map[string]int64 {
	res := map[string]map[string]int64{}
	y := time.Now().Year()
	m := time.Now().Month()
	date := time.Date(y, m, 0, 0, 0, 0, 0, time.UTC)
	maxDay := date.Day()
	for i := 1; i <= maxDay; i++ {
		res[fmt.Sprint(i)] = map[string]int64{}
		for j := 0; j <= 24; j++ {
			res[fmt.Sprint(i)][fmt.Sprint(j)] = 0
		}
	}
	return res
}
