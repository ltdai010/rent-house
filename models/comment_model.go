package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type Comment struct {
	Content   string `json:"content"`
	OwnerID	  string `json:"owner_id"`
	Header    string `json:"header"`
	HouseID	  string `json:"house_id"`
	PostTime  int64  `json:"post_time"`
	Star	  int	 `json:"star"`
	Activate  bool	 `json:"activate"`
}

func (g *Comment) GetCollectionKey() string {
	return consts.COMMENT
}

func (g *Comment) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Comment) GetPaginate(page int, count int) ([]*response.Comment, error) {
	listComment := []*response.Comment{}
	listDoc, err := this.GetCollection().StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q response.Comment
		err = i.DataTo(&q)
		q.CommentID = i.Ref.ID
		listComment = append(listComment, &q)
	}
	return listComment, nil
}

func (this *Comment) PutItem() error {
	res, _, err := client.Collection(this.GetCollectionKey()).Add(ctx, *this)
	if err != nil {
		return err
	}
	_, err = client.Collection(consts.COMMENT_WAIT_LIST).Doc(res.ID).Set(ctx, map[string]string{
		"CommentID" : res.ID,
	})
	return err
}

func (this *Comment) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Comment) DeleteWaitList(id string) error {
	_, err := client.Collection(consts.COMMENT_WAIT_LIST).Doc(id).Delete(ctx)
	return err
}

func (this *Comment) GetFromKey(key string) error {
	doc, err := client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return err
	}
	err = doc.DataTo(this)
	return err
}

func (this *Comment) GetAll() ([]*response.Comment, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listComment := []*response.Comment{}
	for {
		var q response.Comment
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.CommentID = doc.Ref.ID
		listComment = append(listComment, &q)
	}
	return listComment, nil
}

func (this *Comment) GetAllCommentInPost(id string) ([]*response.Comment, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Where("HouseID", "==", id).Documents(ctx)
	listComment := []*response.Comment{}
	for {
		var q response.Comment
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.CommentID = doc.Ref.ID
		listComment = append(listComment, &q)
	}
	return listComment, nil
}

func (this *Comment) GetAllWaitList() ([]string, error) {
	listdoc := client.Collection(consts.COMMENT_WAIT_LIST).Documents(ctx)
	listOwner := []string{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		i, err := doc.DataAt("CommentID")
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, i.(string))
	}
	return listOwner, nil
}

func (this *Comment) GetPaginateWaitList(page int, count int) ([]string, error) {
	listOwner := []string{}
	listDoc, err := client.Collection(consts.COMMENT_WAIT_LIST).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		s, err := i.DataAt("CommentID")
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, s.(string))
	}
	return listOwner, nil
}

func (this *Comment) GetPaginateCommentInPost(id string, page int, count int) ([]*response.Comment, error) {
	listDoc, err := this.GetCollection().Where("PostID", "==", id).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	listComment := []*response.Comment{}
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q response.Comment
		err = i.DataTo(&q)
		q.CommentID = i.Ref.ID
		listComment = append(listComment, &q)
	}
	return listComment, nil
}

func (this *Comment) UpdateItem(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, *this)
	return err
}