package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type Renter struct {
	RenterName 		string 			`json:"renter_name"`
	RenterFullName 	string 		`json:"renter_full_name"`
	Password		string			`json:"password"`
	PhoneNumber		string 			`json:"phone_number"`
	Email			string 			`json:"email"`
	PasswordChanged	int64			`json:"password_changed"`
	ListFavourite	[]string		`json:"list_favourite"`
}

func (g *Renter) GetCollectionKey() string {
	return consts.RENTER
}

func (g *Renter) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}

func (this *Renter) GetPaginate(page int, count int) ([]response.Renter, int, error) {
	listRenter := []response.Renter{}
	start := page * count
	end := start + count
	listDoc, err := this.GetCollection().OrderBy("RenterName", firestore.Asc).StartAt(page*count).Documents(Ctx).GetAll()
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
		var q response.Renter
		err = i.DataTo(&q)
		listRenter = append(listRenter, q)
	}
	return listRenter, len(listDoc), nil
}

func (this *Renter) PutItem() error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(this.RenterName).Set(Ctx, *this)
	return err
}

func (this *Renter) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(Ctx)
	return err
}

func (this *Renter) GetFromKey(key string) error {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(Ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	err = doc.DataTo(this)
	return err
}

func (this *Renter) GetAll() ([]response.Renter, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(Ctx)
	listRenter := []response.Renter{}
	for {
		var q response.Renter
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
		q.RenterID = doc.Ref.ID
		if q.ListFavourite == nil {
			q.ListFavourite = []string{}
		}
		listRenter = append(listRenter, q)
	}
	return listRenter, nil
}

func (this *Renter) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(Ctx, *this)
	return err
}