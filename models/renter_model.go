package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
)

type Renter struct {
	RenterName 		string 		`json:"renter_name"`
	PhoneNumber		string 		`json:"phone_number"`
	Email			string 		`json:"email"`
	ListFavourite	[]string	`json:"list_favourite"`
}

func (g *Renter) GetCollectionKey() string {
	return consts.RENTER
}

func (g *Renter) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Renter) GetPaginate(page int, count int) ([]*Renter, error) {
	listRenter := []*Renter{}
	listDoc, err := this.GetCollection().Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for i := 0; i < page; i++ {
		if len(listDoc) < count {
			return nil, nil
		}
		listDoc, err = this.GetCollection().StartAfter(listDoc[len(listDoc) - 1]).Limit(count).Documents(ctx).GetAll()
		if err != nil {
			return nil, err
		}
	}
	for _, i := range listDoc {
		var q Renter
		err = i.DataTo(&q)
		listRenter = append(listRenter, &q)
	}
	return listRenter, nil
}

func (this *Renter) PutItem() error {
	_, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, this)
	return err
}

func (this *Renter) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Renter) GetFromKey(key string) (*Renter, error) {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = doc.DataTo(this)
	return this, err
}

func (this *Renter) GetAll() ([]*Renter, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listRenter := []*Renter{}
	for {
		var q Renter
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		listRenter = append(listRenter, &q)
	}
	return listRenter, nil
}