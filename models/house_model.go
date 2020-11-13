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


func (g *House) GetCollectionKey() string {
	return consts.HOUSE
}

func (g *House) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *House) GetPaginate(page int, count int) ([]*House, error) {
	listHouse := []*House{}
	listDoc, err := this.GetCollection().StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q House
		err = i.DataTo(&q)
		listHouse = append(listHouse, &q)
	}
	return listHouse, nil
}

func (this *House) PutItem() error {
	//add to collection
	res, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, *this)
	if err != nil {
		return err
	}
	//add to search
	_, err = searchIndex.SaveObject(HouseSearch{
		ObjectID: res.ID,
		House:    *this,
	})
	if err != nil {
		return err
	}
	_, err = client.Collection(consts.HOUSE_WAIT_LIST).Doc(res.ID).Set(ctx, map[string]string{
		"HouseID" : res.ID,
	})
	return err
}

func (this *House) AddToWaitList(time PostTime) error {
	//add to wait list
	_, err := this.GetCollection().Doc(time.HouseID).Set(ctx, time)
	return err
}

func (this *House) Public(time PostTime) error {
	//get item
	doc, err := this.GetCollection().Doc(time.HouseID).Get(ctx)
	if err != nil {
		return err
	}
	h := House{}
	err = doc.DataTo(&h)
	if err != nil {
		return err
	}
	h.PostTime = time.PostTime
	h.ExpiredTime = time.ExpireTime
	h.Activate = true
	_, err = this.GetCollection().Doc(time.HouseID).Set(ctx, h)
	return err
}

func (this *House) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *House) DeleteWaitList(id string) error {
	_, err := client.Collection(consts.HOUSE_WAIT_LIST).Doc(id).Delete(ctx)
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

func (this *House) GetAllHouseOfOwner(id string) ([]*response.House, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Where("OwnerID", "==", id).Documents(ctx)
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

func (this *House) GetPaginateHouseOfUser(id string, page int, count int) ([]*response.House, error) {
	listDoc, err := this.GetCollection().Where("OwnerID", "==", id).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	listComment := []*response.House{}
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q response.House
		err = i.DataTo(&q)
		q.HouseID = i.Ref.ID
		listComment = append(listComment, &q)
	}
	return listComment, nil
}

func (this *House) GetAllWaitList() ([]string, error) {
	listdoc := client.Collection(consts.HOUSE_WAIT_LIST).Documents(ctx)
	listOwner := []string{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		i, err := doc.DataAt("HouseID")
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, i.(string))
	}
	return listOwner, nil
}

func (this *House) GetPaginateWaitList(page int, count int) ([]string, error) {
	listOwner := []string{}
	listDoc, err := client.Collection(consts.HOUSE_WAIT_LIST).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		s, err := i.DataAt("HouseID")
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, s.(string))
	}
	return listOwner, nil
}

func (this *House) UpdateItem(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, *this)
	return err
}