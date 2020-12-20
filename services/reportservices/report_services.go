package reportservices

import (
	"rent-house/models"
	"rent-house/restapi/request"
	"rent-house/restapi/response"
	"time"
)

func AddReport(houseID, renterID string, request request.ReportPost) error {
	//check house exist
	h := &models.House{}
	err := h.GetFromKey(houseID)
	if err != nil {
		return err
	}

	//make report
	rep := &models.Report{
		Title:   request.Title,
		Content:  request.Content,
		RenterID: renterID,
		HouseID:  houseID,
		Seen:     false,
		SendTime: time.Now().Unix(),
	}

	return rep.PutItem()
}

// -1: seen | 0: all | 1: unseen
func GetPageWithFlag(page, count, flag int) ([]response.Report, int, error) {
	rep := &models.Report{}
	if page < 0 || count < 0 {
		return nil, 0, response.BadRequest
	}
	switch flag {
	case -1:
		return rep.GetPageStatus(page, count, true)
	case 0:
		return rep.GetPageAll(page, count)
	case 1:
		return rep.GetPageStatus(page, count, false)
	}
	return nil, 0, response.BadRequest
}

func GetPageInHouse(houseID string, page, count int) ([]response.Report, int, error) {
	rep := &models.Report{}
	if page < 0 || count < 0 {
		return nil, 0, response.BadRequest
	}
	res, total, err :=  rep.GetPageAllInHouse(houseID, page, count)
	if err != nil {
		return nil, 0, err
	}
	go SeenMultiple(res)
	return res, total, nil
}

func SeenMultiple(list []response.Report) {
	for _, i := range list {
		r := &models.Report{}
		err := r.GetFromKey(i.ReportID)
		if err != nil {
			continue
		}
		r.Seen = true
		go r.UpdateItem(i.ReportID)
	}
}

func DeleteReport(id string) error {
	r := &models.Report{}
	return r.Delete(id)
}