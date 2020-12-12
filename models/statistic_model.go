package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
	"strconv"
	"time"
)

type Statistic struct {
	ViewTime       map[string]map[string]int64 `json:"visit_day"`
	ViewLocation   map[string]map[string]int64 `json:"view_location"`
	ViewPriceRange map[string]int64            `json:"view_price_range"`
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
	_, err := Client.Collection(this.GetCollectionKey()).Doc(key).Set(ctx, *this)
	if err != nil {
		return err
	}
	return nil
}

func (this *Statistic) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, *this)
	return err
}

func (this *Statistic) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}


func (this *Statistic) GetFromKey(key string) error {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return err
	}
	err = doc.DataTo(this)
	return err
}


func (this *Statistic) GetAll() ([]Statistic, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(ctx)
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
