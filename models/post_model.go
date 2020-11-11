package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
)

type Post struct {
	HouseID     string `json:"house_id"`
	PostTime	int64  `json:"post_time"`
	Activate	int64  `json:"activate"`
	ExpiredTime	int64  `json:"expired_time"`
}

func (g *Post) GetCollectionKey() string {
	return consts.POST
}

func (g *Post) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Post) GetPaginate(page int, count int) ([]*Post, error) {
	listPost := []*Post{}
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
		var q Post
		err = i.DataTo(&q)
		listPost = append(listPost, &q)
	}
	return listPost, nil
}

func (this *Post) PutItem() error {
	_, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, this)
	return err
}

func (this *Post) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Post) GetFromKey(key string) (*Post, error) {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = doc.DataTo(this)
	return this, err
}

func (this *Post) GetAll() ([]*Post, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listPost := []*Post{}
	for {
		var q Post
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		listPost = append(listPost, &q)
	}
	return listPost, nil
}