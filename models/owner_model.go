package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type Owner struct {
	OwnerName       string  `json:"owner_name"`
	Password        string  `json:"password"`
	OwnerFullName   string  `json:"owner_full_name"`
	Profile         Profile `json:"profile"`
	Address         Address `json:"address"`
	PasswordChanged int64   `json:"password_changed"`
	PostTime        int64   `json:"post_time"`
	Activate        bool    `json:"activate"`
}



func (g *Owner) GetCollectionKey() string {
	return consts.OWNER
}

func (g *Owner) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}

func (this *Owner) GetPaginate(page int, count int) ([]response.Owner, int, error) {
	listOwner := []response.Owner{}
	start := page * count
	end := start + count
	listDoc, err := this.GetCollection().OrderBy("OwnerName", firestore.Asc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	if start > len(listDoc) {
		return nil, 0, response.BadRequest
	}
	if end > len(listDoc) {
		end = len(listDoc)
	}
	for _, i := range listDoc[start : end] {
		var q response.Owner
		err = i.DataTo(&q)
		q.OwnerName = i.Ref.ID
		listOwner = append(listOwner, q)
	}
	return listOwner, len(listDoc), nil
}

func (this *Owner) PutItem() error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(this.OwnerName).Set(Ctx, *this)
	return err
}

func (this *Owner) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(Ctx, *this)
	return err
}

func (this *Owner) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(Ctx)
	return err
}


func (this *Owner) GetFromKey(key string) error {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(Ctx)
	if err != nil {
		return err
	}
	err = doc.DataTo(this)
	return err
}

func (this *Owner) GetAllWaitList() ([]response.Owner, error) {
	listdoc := Client.Collection(consts.OWNER).Where("Activate", "==", false).Documents(Ctx)
	listOwner := []response.Owner{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		o := response.Owner{}
		err = doc.DataTo(&o)
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, o)
	}
	return listOwner, nil
}

func (this *Owner) GetPaginateWaitList(page int, count int) ([]response.Owner, int, error) {
	listOwner := []response.Owner{}
	start := page * count
	end := start + count
	listDoc, err := Client.Collection(consts.OWNER).Where("Activate", "==", false).OrderBy("OwnerName", firestore.Asc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	if start > len(listDoc) {
		return nil, 0, response.BadRequest
	}
	if end > len(listDoc) {
		end = len(listDoc)
	}
	for _, i := range listDoc[start : end]{
		o := response.Owner{}
		err = i.DataTo(&o)
		if err != nil {
			continue
		}
		listOwner = append(listOwner, o)
	}
	return listOwner, len(listDoc), nil
}

func (this *Owner) GetAll() ([]response.Owner, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(Ctx)
	listOwner := []response.Owner{}
	for {
		var q response.Owner
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
		q.OwnerName = doc.Ref.ID
		listOwner = append(listOwner, q)
	}
	return listOwner, nil
}
