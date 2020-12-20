package models

import (
	"cloud.google.com/go/firestore"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type Report struct {
	Tittle   string `json:"tittle"`
	Content  string `json:"content"`
	RenterID string `json:"renter_id"`
	HouseID  string `json:"house_id"`
	Seen	 bool	`json:"seen"`
	SendTime int64  `json:"send_time"`
}

func (g *Report) GetCollectionKey() string {
	return consts.REPORTED
}

func (g *Report) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}

func (g *Report) PutItem() error {
	_, _, err := g.GetCollection().Add(Ctx, g)
	return err
}

func (g *Report) UpdateItem(id string) error {
	_, err := g.GetCollection().Doc(id).Set(Ctx, g)
	return err
}

func (g *Report) GetFromKey(id string) error {
	doc, err := g.GetCollection().Doc(id).Get(Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(g)
}

func (g *Report) GetPageAll(page, count int) ([]response.Report, int, error) {
	res := []response.Report{}
	start := page * count
	end := start + count

	list, err := g.GetCollection().OrderBy("PostTime", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}

	if start > len(list) {
		return nil, 0, response.BadRequest
	}
	if end > len(list) {
		end = len(list)
	}
	for _, i := range list[start : end]{
		r := Report{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		result := ConvertReportResponse(i.Ref.ID, r)
		if result.Tittle == "" {
			continue
		}
		res = append(res, result)
	}
	return res, len(list), nil
}

func (g *Report) GetPageStatus(page, count int, seen bool) ([]response.Report, int, error) {
	res := []response.Report{}
	start := page * count
	end := start + count

	list, err := g.GetCollection().Where("Seen", "==", seen).OrderBy("PostTime", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}

	if start > len(list) {
		return nil, 0, response.BadRequest
	}
	if end > len(list) {
		end = len(list)
	}
	for _, i := range list[start : end] {
		r := Report{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		result := ConvertReportResponse(i.Ref.ID, r)
		if result.Tittle == "" {
			continue
		}
		res = append(res, result)
	}
	return res, len(list), nil
}

func (g *Report) GetPageAllInHouse(houseID string, page, count int) ([]response.Report, int, error) {
	res := []response.Report{}
	start := page * count
	end := start + count

	list, err := g.GetCollection().Where("HouseID", "==", houseID).OrderBy("PostTime", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	if start > len(list) {
		return nil, 0, response.BadRequest
	}
	if end > len(list) {
		end = len(list)
	}

	for _, i := range list[start : end] {
		r := Report{}
		err = i.DataTo(&r)
		if err != nil {
			continue
		}
		result := ConvertReportResponse(i.Ref.ID, r)
		if result.Tittle == "" {
			continue
		}
		res = append(res, result)
	}
	return res, len(list), nil
}

func (r *Report) Delete(id string) error {
	_, err := r.GetCollection().Doc(id).Delete(Ctx)
	return err
}

func ConvertReportResponse(id string, report Report) response.Report {
	house := House{}
	res, err := house.GetResponse(report.HouseID)
	if err != nil {
		return response.Report{}
	}
	return response.Report{
		ReportID: id,
		Tittle:   report.Tittle,
		Content:  report.Content,
		RenterID: report.RenterID,
		House:    res,
		Seen:     report.Seen,
		SendTime: report.SendTime,
	}
}