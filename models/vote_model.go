package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
)

type Vote struct {
	RenterID     string     `json:"renter_id"`
	HouseID		 string		`json:"house_id"`
	Stars     	 int   		`json:"star"`
}

func (g *Vote) GetCollectionKey() string {
	return consts.VOTE
}

func (g *Vote) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Vote) GetPaginate(page int, count int) ([]*Vote, error) {
	listVote := []*Vote{}
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
		var q Vote
		err = i.DataTo(&q)
		listVote = append(listVote, &q)
	}
	return listVote, nil
}

func (this *Vote) PutItem() error {
	_, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, this)
	return err
}

func (this *Vote) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Vote) GetFromKey(key string) (*Vote, error) {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = doc.DataTo(this)
	return this, err
}

func (this *Vote) GetAll() ([]*Vote, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listVote := []*Vote{}
	for {
		var q Vote
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		listVote = append(listVote, &q)
	}
	return listVote, nil
}
