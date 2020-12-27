package statisticservices

import (
	"fmt"
	"rent-house/models"
	"rent-house/restapi/response"
	"time"
)

func MostViewInMonth(length int) ([]response.House, error) {
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
	if m != 12 {
		m++
	} else {
		m = 1
	}
	date := time.Date(y, m, 0, 0, 0, 0, 0, time.UTC)
	maxDay := date.Day()
	for i := 1; i <= maxDay; i++ {
		res[fmt.Sprint(i)] = map[string]int64{}
		for j := 0; j < 24; j++ {
			res[fmt.Sprint(i)][fmt.Sprint(j)] = 0
		}
	}
	return res
}

func HouseInLocation(length int) ([]models.InLocation, error) {
	statistic := &models.Statistic{}
	result, err := statistic.GetNumberHouseInLocation()
	return FindMaxInLocation(result, length), err
}

func ViewInLocation(length int) ([]models.InLocation, error) {
	statistic := &models.Statistic{}
	err := statistic.GetFromKey(statistic.GetKeyNow())
	if err != nil {
		return []models.InLocation{}, err
	}
	return FindMaxInLocation(statistic.ViewLocation, length), nil
}

func FindMaxInLocation(mapView map[string]map[string]int64, length int) []models.InLocation {
	res := []models.InLocation{}
	for province, provinceView := range mapView {
		for district, districtView := range provinceView {
			i := 0
			l := length
			if len(res) < length {
				l = len(res)
			}
			for i = 0; i < len(res); i++ {
				if districtView >= res[i].Number {
					pros := res[0:i]
					cons := make([]models.InLocation, len(res[i:l]))
					copy(cons, res[i:l])
					res = append(pros, models.InLocation{
						Number:   districtView,
						Location: district + " - " + province,
					})
					res = append(res, cons...)
					break
				}
			}
			if i == len(res) && len(res) < length{
				res = append(res, models.InLocation{
					Number:   districtView,
					Location: district + " - " + province,
				})
			}
		}
	}
	return res
}

func ViewByPrice() (map[string]int64, error) {
	statistic := &models.Statistic{}
	err := statistic.GetFromKey(statistic.GetKeyNow())
	if err != nil {
		return map[string]int64{}, err
	}
	return statistic.ViewPriceRange, nil
}
