package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
)

type Comment struct {
	Content   string `json:"content"`
	Header    string `json:"header"`
	PostTime  int64  `json:"post_time"`
	Activate  bool	 `json:"activate"`
}

func (g *Comment) GetCollectionKey() string {
	return consts.COMMENT
}

func (g *Comment) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Comment) GetPaginate(page int, count int) ([]*Comment, error) {
	listComment := []*Comment{}
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
		var q Comment
		err = i.DataTo(&q)
		listComment = append(listComment, &q)
	}
	return listComment, nil
}

func (this *Comment) PutItem() error {
	_, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, this)
	return err
}

func (this *Comment) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Comment) GetFromKey(key string) (*Comment, error) {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = doc.DataTo(this)
	return this, err
}

func (this *Comment) GetAll() ([]*Comment, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listComment := []*Comment{}
	for {
		var q Comment
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		listComment = append(listComment, &q)
	}
	return listComment, nil
}

