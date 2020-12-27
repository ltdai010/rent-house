package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
	"rent-house/consts"
	"strconv"
	"time"
)

type Statistic struct {
	ViewTime       map[string]map[string]int64 `json:"visit_day"`
	ViewLocation   map[string]map[string]int64 `json:"view_location"`
	ViewPriceRange map[string]int64            `json:"view_price_range"`
}

type InLocation struct {
	Number	int64	`json:"number"`
	Location string `json:"location"`
}

func (g *Statistic) GetKeyNow() string {
	return strconv.Itoa(time.Now().Year()) + "-" + time.Now().Month().String()
}

func (g *Statistic) GetCollectionKey() string {
	return consts.STATISTIC
}

func (g *Statistic) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}


func (this *Statistic) PutItem() error {
	key := strconv.Itoa(time.Now().Year()) + "-" + time.Now().Month().String()
	_, err := Client.Collection(this.GetCollectionKey()).Doc(key).Set(Ctx, *this)
	if err != nil {
		return err
	}
	return nil
}

func (this *Statistic) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(Ctx, *this)
	return err
}

func (this *Statistic) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(Ctx)
	return err
}


func (this *Statistic) GetFromKey(key string) error {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(Ctx)
	if err != nil {
		return err
	}
	err = doc.DataTo(this)
	return err
}

func (this *Statistic) DecreaseHouseInDistrict(province, district string) {
	res, err := this.GetCollection().Doc("location").Get(Ctx)
	if err != nil {
		log.Println(err, "   models/statistic_model.go:68")
		return
	}
	data := map[string]map[string]int64{}
	err = res.DataTo(&data)
	if err != nil {
		log.Println(err, "   models/statistic_model.go:74")
		return
	}
	if data[province] == nil {
		return
	}
	if data[province][district] == 0 {
		return
	}
	data[province][district]--
	_, err = this.GetCollection().Doc("location").Set(Ctx, data)
	if err != nil {
		log.Println(err, "   models/statistic_model.go:85")
	}
}

func (this *Statistic) IncreaseHouseInDistrict(province, district string) {
	res, err := this.GetCollection().Doc("location").Get(Ctx)
	if err != nil {
		log.Println(err, "   models/statistic_model.go:93")
		return
	}
	data := map[string]map[string]int64{}
	err = res.DataTo(&data)
	if err != nil {
		log.Println(err, "   models/statistic_model.go:99")
		return
	}
	if data[province] == nil {
		data[province] = map[string]int64{}
	}
	data[province][district]++
	_, err = this.GetCollection().Doc("location").Set(Ctx, data)
	if err != nil {
		log.Println(err, "   models/statistic_model.go:108")
	}
}

func (this *Statistic) GetNumberHouseInLocation() (map[string]map[string]int64, error) {
	res := map[string]map[string]int64{}
	doc, err := Client.Collection(consts.STATISTIC).Doc("location").Get(Ctx)
	if err != nil {
		return nil, err
	}
	err = doc.DataTo(&res)
	if err != nil {
		return map[string]map[string]int64{}, err
	}
	return res, nil
}

func (this *Statistic) CalculateHouseInLocation() {
	res := map[string]map[string]int64{}
	listAll, err := Client.Collection(consts.HOUSE).Documents(Ctx).GetAll()
	for _, i := range listAll {
		h := House{}
		err = i.DataTo(&h)
		if err != nil {
			continue
		}
		a := &Address{}
		err = a.FindAddress(h.CommuneCode)
		if err != nil {
			continue
		}
		if res[a.Province] == nil {
			res[a.Province] = map[string]int64{}
		}
		res[a.Province][a.District]++
	}
	this.GetCollection().Doc("location").Set(Ctx, res)
}

func (this *Statistic) CalculateViewInLocation() {
	list, _ := Client.Collection(consts.HOUSE).Documents(Ctx).GetAll()
	makeMap := map[string]map[string]int64{}
	for _, i := range list {
		h := House{}
		err := i.DataTo(&h)
		if err != nil {
			continue
		}
		a := &Address{}
		a.FindAddress(h.CommuneCode)
		if makeMap[a.Province] == nil {
			makeMap[a.Province] = map[string]int64{}
		}
		makeMap[a.Province][a.District]++
	}
	stat := &Statistic{}
	stat.GetFromKey(this.GetKeyNow())
	stat.ViewLocation = makeMap
	stat.PutItem()
}

func (this *Statistic) GetAll() ([]Statistic, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(Ctx)
	list := []Statistic{}
	for {
		var q Statistic
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		list = append(list, q)
	}
	return list, nil
}
