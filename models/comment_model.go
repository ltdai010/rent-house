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

func (this *Comment) GetPaginate(page int, count int) ([]response.Comment, int, error) {
	listComment := []response.Comment{}
	listDoc, err := this.GetCollection().Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(listDoc)
	if page * count > total {
		return nil, 0, response.BadRequest
	}
	end := (page + 1) * count
	if end > total {
		end = total
	}
	for _, i := range listDoc[page * count : end]{
		var q Comment
		err = i.DataTo(&q)
		if err != nil {
			continue
		}
		result := ConvertCommentResponse(i.Ref.ID, q)
		listComment = append(listComment, result)
	}
	return listComment, total, nil
}

func (this *Comment) PutItem() error {
	_, _, err := Client.Collection(this.GetCollectionKey()).Add(Ctx, *this)
	return err
}

func (this *Comment) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(Ctx)
	return err
}

func (this *Comment) GetFromKey(key string) error {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(Ctx)
	if err != nil {
		return err
	}
	err = doc.DataTo(this)
	return err
}

func (this *Comment) GetAll() ([]response.Comment, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(Ctx)
	listComment := []response.Comment{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var q Comment
		err = doc.DataTo(&q)
		if err != nil {
			continue
		}
		result := ConvertCommentResponse(doc.Ref.ID, q)
		if result.Content == "" {
			continue
		}
		listComment = append(listComment, result)
	}
	return listComment, nil
}

func (this *Comment) GetAllCommentActiveInHouse(id string) ([]response.Comment, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Where("HouseID", "==", id).Documents(Ctx)
	listComment := []response.Comment{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var q Comment
		err = doc.DataTo(&q)
		if err != nil {
			continue
		}
		if q.Activate == false {
			continue
		}
		result := ConvertCommentResponse(doc.Ref.ID, q)
		if result.Content == "" {
			continue
		}
		listComment = append(listComment, result)
	}
	return listComment, nil
}

func (this *Comment) GetAllCommentInHouse(id string) ([]response.Comment, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Where("HouseID", "==", id).Documents(Ctx)
	listComment := []response.Comment{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var q Comment
		err = doc.DataTo(&q)
		if err != nil {
			continue
		}
		result := ConvertCommentResponse(doc.Ref.ID, q)
		if result.Content == "" {
			continue
		}
		listComment = append(listComment, result)
	}
	return listComment, nil
}

func (this *Comment) GetPaginateCommentInHouse( id string, page, count int) ([]response.Comment, int, error) {
	listdoc, err := Client.Collection(this.GetCollectionKey()).Where("HouseID", "==", id).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(listdoc)
	if page * count > total {
		return nil, 0, response.BadRequest
	}
	end := (page + 1) * count
	if end > total {
		end = total
	}
	listComment := []response.Comment{}
	for _, i := range listdoc[page * count : end]{
		var q Comment
		err = i.DataTo(&q)
		if err != nil {
			continue
		}
		result := ConvertCommentResponse(i.Ref.ID, q)
		if result.Content == "" {
			continue
		}
		listComment = append(listComment, result)
	}
	return listComment, total, nil
}

func (this *Comment) GetAllWaitList() ([]response.Comment, error) {
	listdoc := Client.Collection(consts.COMMENT).Where("Activate", "==", false).Documents(Ctx)
	listComment := []response.Comment{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var q Comment
		err = doc.DataTo(&q)
		if err != nil {
			continue
		}
		result := ConvertCommentResponse(doc.Ref.ID, q)
		if result.Content == "" {
			continue
		}
		listComment = append(listComment, result)
	}
	return listComment, nil
}

func (this *Comment) GetPaginateWaitList(page int, count int) ([]response.Comment, int, error) {
	listDoc, err := Client.Collection(consts.COMMENT).Where("Activate", "==", false).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(listDoc)
	if page * count > total {
		return nil, 0,  response.BadRequest
	}
	end := (page + 1) * count
	if end > total {
		end = total
	}
	listComment := []response.Comment{}
	for _, i := range listDoc[page * count : end]{
		var q Comment
		err = i.DataTo(&q)
		if err != nil {
			continue
		}
		result := ConvertCommentResponse(i.Ref.ID, q)
		if result.Content == "" {
			continue
		}
		listComment = append(listComment, result)
	}
	return listComment, total, nil
}

func (this *Comment) GetPaginateCommentActiveInHouse(id string, page int, count int) ([]response.Comment, int, error) {
	list, err := this.GetCollection().Where("Activate", "==", true).Where("HouseID", "==", id).Documents(Ctx).GetAll()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	total := len(list)
	if page * count > total {
		return nil, 0, response.BadRequest
	}
	end := (page + 1) * count
	if end > total {
		end = total
	}
	listComment := []response.Comment{}
	for _, i := range list[page * count : end]{
		var q Comment
		err = i.DataTo(&q)
		if err != nil {
			continue
		}
		result := ConvertCommentResponse(i.Ref.ID, q)
		if result.Content == "" {
			continue
		}
		listComment = append(listComment, result)
	}
	return listComment, total, nil
}

func (this *Comment) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(Ctx, *this)
	return err
}

func ConvertCommentResponse(id string, comment Comment) response.Comment {
	renter := &Renter{}
	err := renter.GetFromKey(comment.RenterID)
	if err != nil {
		return response.Comment{}
	}
	return response.Comment{
		CommentID:  id,
		Content:    comment.Content,
		RenterName: renter.RenterFullName,
		HouseID:    comment.HouseID,
		PostTime:   comment.PostTime,
		Star:       comment.Star,
		Activate:   comment.Activate,
	}
}