package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type House struct {
	OwnerID        string         `json:"owner_id"`
	HouseType      HouseType      `json:"house_type"`
	PricePerMonth  int            `json:"price"`
	PricePerYear   int            `json:"price_per_year"`
	Address        Address        `json:"address"`
	Infrastructure Infrastructure `json:"infrastructure"`
	NearBy         []string       `json:"near_by"`
	WithOwner      bool           `json:"with_owner"`
	ImageLink      []string       `json:"image_link"`
	Header         string         `json:"header"`
	Content        string         `json:"content"`
	PostTime	   int64  		  `json:"post_time"`
	Activate	   bool  		  `json:"activate"`
	ExpiredTime	   int64  		  `json:"expired_time"`
}

type HouseSearch struct {
	ObjectID string `json:"objectID"`
	House
}

type Address struct {
	Province string `json:"province"`
	District string `json:"district"`
	Street   string `json:"street"`
}

type Infrastructure struct {
	PrivateBathroom bool   `json:"private_bathroom"`
	Heater          bool   `json:"heater"`
	AirCondition    bool   `json:"air_condition"`
	Balcony         bool   `json:"balcony"`
	ElectricPrice   int    `json:"electric_price"`
	WaterPrice      int    `json:"water_price"`
	Other           string `json:"other"`
}

type HouseType int

const (
	Room = iota
	MiniApartment
	FullHouse
	Apartment
)

func (g *House) GetCollectionKey() string {
	return consts.HOUSE
}

func (g *House) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *House) GetPaginate(page int, count int) ([]*House, error) {
	listHouse := []*House{}
	listDoc, err := this.GetCollection().Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for i := 0; i < page; i++ {
		if len(listDoc) < count {
			return nil, nil
		}
		listDoc, err = this.GetCollection().StartAfter(listDoc[len(listDoc)-1]).Limit(count).Documents(ctx).GetAll()
		if err != nil {
			return nil, err
		}
	}
	for _, i := range listDoc {
		var q House
		err = i.DataTo(&q)
		listHouse = append(listHouse, &q)
	}
	return listHouse, nil
}

func (this *House) PutItem() error {
	res, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, this)
	if err != nil {
		return err
	}
	_, err = searchIndex.SaveObject(HouseSearch{
		ObjectID: res.ID,
		House:    *this,
	})
	return err
}

func (this *House) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *House) GetFromKey(key string) (*House, error) {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = doc.DataTo(this)
	return this, err
}

func (this *House) GetAll() ([]*response.House, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listHouse := []*response.House{}
	for {
		var q response.House
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.HouseID = doc.Ref.ID
		listHouse = append(listHouse, &q)
	}
	return listHouse, nil
}

func (this *House) UpdateItem(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, this)
	return err
}