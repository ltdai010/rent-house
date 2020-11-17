package models

import (
	"cloud.google.com/go/firestore"
	"errors"
	"google.golang.org/api/iterator"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type Renter struct {
	RenterName 		string 			`json:"renter_name"`
	RenterFullName 		string 		`json:"renter_full_name"`
	Password		string			`json:"password"`
	PhoneNumber		string 			`json:"phone_number"`
	Email			string 			`json:"email"`
	ListFavourite	[]string		`json:"list_favourite"`
}

func (g *Renter) GetCollectionKey() string {
	return consts.RENTER
}

func (g *Renter) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Renter) GetPaginate(page int, count int) ([]*Renter, error) {
	listRenter := []*Renter{}
	listDoc, err := this.GetCollection().StartAt(page*count).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q Renter
		err = i.DataTo(&q)
		listRenter = append(listRenter, &q)
	}
	return listRenter, nil
}

func (this *Renter) PutItem() error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(this.RenterName).Set(ctx, *this)
	return err
}

func (this *Renter) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Renter) GetFromKey(key string) error {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return err
	}
	if doc == nil {
		return errors.New("not exist" + this.GetCollectionKey())
	}
	err = doc.DataTo(this)
	return err
}

func (this *Renter) GetAll() ([]*response.Renter, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listRenter := []*response.Renter{}
	for {
		var q response.Renter
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.RenterID = doc.Ref.ID
		listRenter = append(listRenter, &q)
	}
	return listRenter, nil
}

func (this *Renter) UpdateItem(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, *this)
	return err
}