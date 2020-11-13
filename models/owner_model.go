package models

import (
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"rent-house/consts"
	"rent-house/restapi/response"
)

type Owner struct {
	OwnerName 		string  	`json:"owner_name"`
	Password		string		`json:"password"`
	OwnerFullName	string		`json:"owner_full_name"`
	Profile   		Profile 	`json:"profile"`
	Address   		Address 	`json:"address"`
	Activate  		bool	 	`json:"activate"`
}

type OwnerLogin struct {
	OwnerName 		string  	`json:"owner_name"`
	Password		string		`json:"password"`
}

func (this OwnerLogin) GetPassword() string {
	return this.Password
}

func (this OwnerLogin) GetUsername() string {
	return this.OwnerName
}

func (g *Owner) GetCollectionKey() string {
	return consts.OWNER
}

func (g *Owner) GetCollection() *firestore.CollectionRef {
	return client.Collection(g.GetCollectionKey())
}

func (this *Owner) GetPaginate(page int, count int) ([]*Owner, error) {
	listOwner := []*Owner{}
	listDoc, err := this.GetCollection().StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q Owner
		err = i.DataTo(&q)
		listOwner = append(listOwner, &q)
	}
	return listOwner, nil
}

func (this *Owner) PutItem() error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(this.OwnerName).Set(ctx, *this)
	if err != nil {
		return err
	}
	_, err = client.Collection(consts.OWNER_WAIT_LIST).Doc(this.OwnerName).Set(ctx, map[string]string{
		"OwnerName" : this.OwnerName,
	})
	return err
}

func (this *Owner) UpdateItem(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, *this)
	return err
}

func (this *Owner) Delete(id string) error {
	_, err := client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	return err
}

func (this *Owner) DeleteWaitList(id string) error {
	_, err := client.Collection(consts.OWNER_WAIT_LIST).Doc(id).Delete(ctx)
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

func (this *Owner) GetAllWaitList() ([]string, error) {
	listdoc := client.Collection(consts.OWNER_WAIT_LIST).Documents(ctx)
	listOwner := []string{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		i, err := doc.DataAt("OwnerID")
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, i.(string))
	}
	return listOwner, nil
}

func (this *Owner) GetPaginateWaitList(page int, count int) ([]string, error) {
	listOwner := []string{}
	listDoc, err := client.Collection(consts.OWNER_WAIT_LIST).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		s, err := i.DataAt("OwnerID")
		if err != nil {
			return nil, err
		}
		listOwner = append(listOwner, s.(string))
	}
	return listOwner, nil
}

func (this *Owner) GetAll() ([]*response.Owner, error) {
	listdoc := client.Collection(this.GetCollectionKey()).Documents(ctx)
	listOwner := []*response.Owner{}
	for {
		var q response.Owner
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.OwnerID = doc.Ref.ID
		listOwner = append(listOwner, &q)
	}
	return listOwner, nil
}
