package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
)

type Owner struct {
	OwnerName string  `json:"owner_name"`
	Profile   Profile `json:"profile"`
	Address   Address `json:"address"`
	Activate  bool	  `json:"activate"`
}

type Profile struct {
	IDCard      string `json:"id_card"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (g *Owner) GetCollectionKey() string {
	return consts.OWNER
}

func (g *Owner) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Owner) GetPaginate(page int, count int) ([]*Owner, error) {
	listOwner := []*Owner{}
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
		var q Owner
		err = i.DataTo(&q)
		listOwner = append(listOwner, &q)
	}
	return listOwner, nil
}

func (this *Owner) PutItem() error {
	_, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, this)
	return err
}

func (this *Owner) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Owner) GetFromKey(key string) (*Owner, error) {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = doc.DataTo(this)
	return this, err
}

func (this *Owner) GetAll() ([]*Owner, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listOwner := []*Owner{}
	for {
		var q Owner
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, &q)
	}
	return listOwner, nil
}
