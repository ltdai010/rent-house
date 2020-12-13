package models

import (
	"cloud.google.com/go/firestore"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"mime/multipart"
	"rent-house/consts"
	"rent-house/restapi/response"
	"time"
)

type House struct {
	OwnerID        string         `json:"owner_id"`
	HouseType      HouseType      `json:"house_type"`
	Price		   float64     	  `json:"price"`
	Unit 		   Unit			  `json:"unit"`
	Address        Address        `json:"address"`
	CommuneCode	   string		  `json:"commune_code"`
	Infrastructure Infrastructure `json:"infrastructure"`
	NearBy         []string       `json:"near_by"`
	PreOrder	   int			  `json:"pre_order"`
	Surface		   int			  `json:"surface"`
	WithOwner      bool           `json:"with_owner"`
	ImageLink      []string       `json:"image_link"`
	LastViewed	   int64 		  `json:"last_viewed"`
	MonthlyView	   int			  `json:"monthly_view"`
	Header         string         `json:"header"`
	View		   int64 		  `json:"view"`
	Like		   int64		  `json:"like"`
	Rented		   bool			  `json:"rented"`
	Content        string         `json:"content"`
	PostTime	   int64  		  `json:"post_time"`
	Status	       Status  		  `json:"status"`
	Review         map[string]int `json:"review"`
	AppearTime	   int64		  `json:"appear_time"`
	ExpiredTime	   int64  		  `json:"expired_time"`
	AdminComment   string		  `json:"admin_comment"`
}
type HouseSearch struct {
	ObjectID 	   string 		  `json:"objectID"`
	OwnerID        string         `json:"owner_id"`
	NearBy         []string       `json:"near_by"`
	Header         string         `json:"header"`
	Content        string         `json:"content"`
}

func (g *House) GetCollectionKey() string {
	return consts.HOUSE
}

func (g *House) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}

func (this *House) GetMaxViewHouseInMonth() (response.House, error) {
	ref := Client.Collection(consts.HOUSE).OrderBy("MonthlyView", firestore.Asc).Limit(1).Documents(ctx)
	doc, err := ref.Next()
	if err != nil {
		return response.House{}, err
	}
	res := response.House{}
	err = doc.DataTo(&res)
	if err != nil {
		return response.House{}, err
	}
	res.HouseID = doc.Ref.ID
	return res, nil
}

func (this *House) FindMaxViewHouse() (response.House, error) {
	ref := Client.Collection(consts.HOUSE).OrderBy("View", firestore.Asc).Limit(1).Documents(ctx)
	doc, err := ref.Next()
	if err != nil {
		return response.House{}, err
	}
	res := response.House{}
	err = doc.DataTo(&res)
	res.HouseID = doc.Ref.ID
	return response.House{}, err
}

func (this *House) FindMaxLikeHouse() (response.House, error) {
	ref := Client.Collection(consts.HOUSE).OrderBy("Like", firestore.Asc).Limit(1).Documents(ctx)
	doc, err := ref.Next()
	if err != nil {
		return response.House{}, err
	}
	res := response.House{}
	err = doc.DataTo(&res)
	res.HouseID = doc.Ref.ID
	return response.House{}, err
}

func (this *House) GetPaginate(page int, count int) ([]*House, error) {
	listHouse := []*House{}
	listDoc, err := this.GetCollection().OrderBy("PostTime", firestore.Asc).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q House
		err = i.DataTo(&q)
		listHouse = append(listHouse, &q)
	}
	return listHouse, nil
}

func (this *House) PutItem() (string, error) {
	//add to collection
	res, _, err := Client.Collection(this.GetCollectionKey()).Add(ctx, *this)
	if err != nil {
		return "", err
	}
	//add to search
	go searchIndex.SaveObject(HouseSearch{
		ObjectID:       res.ID,
		OwnerID:        this.OwnerID,
		NearBy:         this.NearBy,
		Header:         this.Header,
		Content:        this.Content,
	})
	return res.ID, err
}

func (this *House) Public(time PostTime) error {
	//get item
	doc, err := this.GetCollection().Doc(time.HouseID).Get(ctx)
	if err != nil {
		return err
	}
	h := House{}
	err = doc.DataTo(&h)
	if err != nil {
		return err
	}
	h.PostTime = time.PostTime
	h.ExpiredTime = time.ExpireTime
	h.Status = Activated
	_, err = this.GetCollection().Doc(time.HouseID).Set(ctx, h)
	return err
}

func (this *House) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(ctx)
	if err != nil {
		return err
	}
	_, err = searchIndex.Delete(id)
	return err
}

func (this *House) AddImage(file multipart.File) (string, error) {
	ref, _, err := Client.Collection(consts.IMAGE_LINK).Add(ctx, map[string]string{
		"Image" : "link",
	})
	if err != nil {
		return "", err
	}
	wc := bucket.Object(ref.ID).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	return ref.ID, nil
}

func (this *House) GetFromKey(key string) (error) {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *House) GetResponse(key string) (response.House, error) {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(ctx)
	if err != nil {
		return response.House{}, err
	}
	res := response.House{}
	res.HouseID = doc.Ref.ID
	err = doc.DataTo(&res)
	return res, err
}

func (this *House) GetAllActivate() ([]response.House, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Where("ExpiredTime", ">", time.Now().Unix()).Documents(ctx)
	listHouse := []response.House{}
	for {
		var q response.House
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
		q.HouseID = doc.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, nil
}

func (this *House) GetPageActivate(page, count int) ([]response.House, error) {
	listdoc, err := Client.Collection(this.GetCollectionKey()).OrderBy("PostTime", firestore.Asc).StartAfter(page * count).Limit(count).Documents(ctx).GetAll()
	listHouse := []response.House{}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	now := time.Now().Unix()
	for _, i := range listdoc {
		var q response.House
		err = i.DataTo(&q)
		q.HouseID = i.Ref.ID
		if q.ExpiredTime > now {
			listHouse = append(listHouse, q)
		}
	}
	return listHouse, nil
}

func (this *House) GetAll() ([]response.House, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(ctx)
	listHouse := []response.House{}
	for {
		var q response.House
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		err = doc.DataTo(&q)
		if err != nil {
			return nil, err
		}
		q.HouseID = doc.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, nil
}

func (this *House) GetAllHouseOfOwner(id string) ([]response.House, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Where("OwnerID", "==", id).Documents(ctx)
	listHouse := []response.House{}
	for {
		var q response.House
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
		q.HouseID = doc.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, nil
}

func (this *House) GetPaginateHouseOfUser(id string, page int, count int) ([]response.House, error) {
	listDoc, err := this.GetCollection().Where("OwnerID", "==", id).OrderBy("PostTime", firestore.Asc).StartAt(page * count).Limit(count).Documents(ctx).GetAll()
	listHouse := []response.House{}
	if err != nil {
		log.Println(err)
		return listHouse, err
	}
	for _, i := range listDoc {
		var q response.House
		err = i.DataTo(&q)
		q.HouseID = i.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, nil
}

func (this *House) GetAllByStatus(status Status) ([]response.House, error) {
	listdoc := Client.Collection(consts.HOUSE).Where("Status", "==", status).Documents(ctx)
	listHouse := []response.House{}
	for {
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		h := response.House{}
		err = doc.DataTo(&h)
		if err != nil {
			continue
		}
		h.HouseID = doc.Ref.ID
		listHouse = append(listHouse, h)
	}
	return listHouse, nil
}

func (this *House) GetPaginateByStatus(status Status, page int, count int) ([]response.House, error) {
	listHouse := []response.House{}
	listDoc := Client.Collection(consts.HOUSE).Where("Status", "==", status).OrderBy("PostTime", firestore.Asc).StartAt(page * count).Limit(count).Documents(ctx)
	for  {
		doc, err := listDoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		h := response.House{}
		err = doc.DataTo(&h)
		if err != nil {
			continue
		}
		h.HouseID = doc.Ref.ID
		listHouse = append(listHouse, h)
	}
	return listHouse, nil
}

func (this *House) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(ctx, *this)
	return err
}

func (this *House) SearchAllItem(key string) ([]response.House, error) {
	res, err := searchIndex.Search(key)
	if err != nil {
		return []response.House{}, err
	}
	list := []response.House{}
	results := []HouseSearch{}
	err = res.UnmarshalHits(&results)
	if err != nil {
		return []response.House{}, err
	}
	now := time.Now().Unix()
	for _, i := range results {
		h := &House{}
		resH, err := h.GetResponse(i.ObjectID)
		if err != nil {
			return []response.House{}, err
		}
		if h.ExpiredTime > now {
			list = append(list, resH)
		}
	}
	return list, nil
}

func (this *House) SearchPaginateItem(key string, page, count int) ([]response.House, error) {
	res, err := searchIndex.Search(key, opt.Page(page), opt.HitsPerPage(count))
	if err != nil {
		return []response.House{}, err
	}
	list := []response.House{}
	results := []HouseSearch{}
	err = res.UnmarshalHits(&results)
	if err != nil {
		return []response.House{}, err
	}
	now := time.Now().Unix()
	for _, i := range results {
		h := &House{}
		resH, err := h.GetResponse(i.ObjectID)
		if err != nil {
			return []response.House{}, err
		}
		if h.ExpiredTime > now {
			list = append(list, resH)
		}
	}
	return list, nil
}