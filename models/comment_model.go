package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type Comment struct {
	Content  string `json:"content"`
	RenterID string `json:"renter_id"`
	Header   string `json:"header"`
	HouseID  string `json:"house_id"`
	PostTime int64  `json:"post_time"`
	Star     int    `json:"star"`
	Activate bool   `json:"activate"`
}

func (g *Comment) GetCollectionKey() string {
	return consts.COMMENT
}

func (g *Comment) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}

func (this *Comment) GetPaginate(page int, count int) ([]response.Comment, error) {
	listComment := []response.Comment{}
	listDoc, err := this.GetCollection().OrderBy("PostTime", firestore.Asc).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q response.Comment
		err = i.DataTo(&q)
		q.CommentID = i.Ref.ID
		listComment = append(listComment, q)
	}
	return listComment, nil
}

func (this *Comment) PutItem() error {
	_, _, err := Client.Collection(this.GetCollectionKey()).Add(ctx, *this)
	return err
}

func (this *Comment) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Comment) GetFromKey(key string) error {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return err
	}
	err = doc.DataTo(this)
	return err
}

func (this *Comment) GetAll() ([]response.Comment, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(ctx)
	listComment := []response.Comment{}
	for {
		var q response.Comment
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
		q.CommentID = doc.Ref.ID
		listComment = append(listComment, q)
	}
	return listComment, nil
}

func (this *Comment) GetAllCommentActiveInHouse(id string) ([]response.Comment, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Where("HouseID", "==", id).Documents(ctx)
	listComment := []response.Comment{}
	for {
		var q response.Comment
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
		if q.Activate == false {
			continue
		}
		q.CommentID = doc.Ref.ID
		listComment = append(listComment, q)
	}
	return listComment, nil
}

func (this *Comment) GetAllCommentInHouse(id string) ([]response.Comment, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Where("HouseID", "==", id).Documents(ctx)
	listComment := []response.Comment{}
	for {
		var q response.Comment
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
		q.CommentID = doc.Ref.ID
		listComment = append(listComment, q)
	}
	return listComment, nil
}

func (this *Comment) GetAllWaitList() ([]response.Comment, error) {
	listdoc := Client.Collection(consts.COMMENT).Where("Activate", "==", false).Documents(ctx)
	listComment := []response.Comment{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		c := response.Comment{}
		err = doc.DataTo(&c)
		if err != nil {
			return nil, err
		}
		c.CommentID = doc.Ref.ID
		listComment = append(listComment, c)
	}
	return listComment, nil
}

func (this *Comment) GetPaginateWaitList(page int, count int) ([]response.Comment, error) {
	listDoc := Client.Collection(consts.COMMENT).Where("Activate", "==", false).OrderBy("PostTime", firestore.Asc).StartAt(page * count).Limit(count).Documents(ctx)
	listComment := []response.Comment{}
	for {
		doc, err := listDoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		c := response.Comment{}
		err = doc.DataTo(&c)
		if err != nil {
			return nil, err
		}
		c.CommentID = doc.Ref.ID
		listComment = append(listComment, c)
	}
	return listComment, nil
}

func (this *Comment) GetPaginateCommentActiveInHouse(id string, page int, count int) ([]response.Comment, error) {
	listDoc := this.GetCollection().Where("HouseID", "==", id).OrderBy("PostTime", firestore.Asc).StartAt(page * count).Limit(count).Documents(ctx)
	listComment := []response.Comment{}
	for {
		doc, err := listDoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			return nil, err
		}
		c := response.Comment{}
		err = doc.DataTo(&c)
		if err != nil {
			return nil, err
		}
		if !c.Activate {
			continue
		}
		c.CommentID = doc.Ref.ID
		listComment = append(listComment, c)
	}
	return listComment, nil
}

func (this *Comment) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, *this)
	return err
}