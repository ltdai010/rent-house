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
	"strconv"
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
	NearBy         []string       `json:"near_by"`
	Header         string         `json:"header"`
	Price          float64		  `json:"price"`
}

func (g *House) GetCollectionKey() string {
	return consts.HOUSE
}

func (g *House) GetCollection() *firestore.CollectionRef {
	return Client.Collection(g.GetCollectionKey())
}

func (this *House) GetMaxViewHouseInMonth(length int) ([]response.House, error) {
	res := []response.House{}
	ref, err := Client.Collection(consts.HOUSE).OrderBy("MonthlyView", firestore.Desc).Limit(length).Documents(Ctx).GetAll()
	for _, i := range ref {
		h := response.House{}
		err = i.DataTo(&h)
		if err != nil {
			continue
		}
		h.HouseID = i.Ref.ID
		res = append(res, h)
	}
	return res, nil
}

func (this *House) GetActiveHouseByListID(ids []string) ([]response.House, error) {
	res := []response.House{}
	for _,i := range ids {
		h := response.House{}
		doc, err := this.GetCollection().Doc(i).Get(Ctx)
		if err != nil {
			continue
		}
		err = doc.DataTo(&h)
		if err != nil {
			continue
		}
		if h.ExpiredTime < time.Now().Unix() {
			continue
		}
		h.HouseID = doc.Ref.ID
		res = append(res, h)
	}
	return res, nil
}

func (this *House) FindMaxViewHouse() (response.House, error) {
	ref := Client.Collection(consts.HOUSE).OrderBy("View", firestore.Desc).Limit(1).Documents(Ctx)
	doc, err := ref.Next()
	if err != nil {
		return response.House{}, err
	}
	res := response.House{}
	err = doc.DataTo(&res)
	res.HouseID = doc.Ref.ID
	return response.House{}, err
}

func (this *House) GetPaginate(page int, count int) ([]response.House, int, error) {
	listHouse := []response.House{}
	start := page * count
	end := start + count

	listDoc, err := this.GetCollection().OrderBy("PostTime", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}

	if start > len(listDoc) {
		return nil, 0 , response.BadRequest
	}
	if end > len(listDoc) {
		end = len(listDoc)
	}
	for _, i := range listDoc[start : end] {
		var q response.House
		err = i.DataTo(&q)
		q.HouseID = i.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, len(listDoc), nil
}

func (this *House) GetPaginateByLike(page int, count int) ([]response.House, int, error) {
	listHouse := []response.House{}
	listDoc, err := this.GetCollection().OrderBy("Like", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for _, i := range listDoc {
		var q response.House
		err = i.DataTo(&q)
		if q.ExpiredTime < time.Now().Unix() {
			continue
		}
		q.HouseID = i.Ref.ID
		listHouse = append(listHouse, q)
	}
	total := len(listHouse)
	if page * count > total {
		return nil, 0, response.BadRequest
	}
	end := (page + 1) * count
	if end > total {
		end = total
	}
	return listHouse[page * count:end], total, nil
}

func (this *House) GetAllByLike() ([]response.House, error) {
	listHouse := []response.House{}
	listDoc, err := this.GetCollection().Where("ExpiredTime", ">", time.Now().Unix()).OrderBy("ExpiredTime", firestore.Asc).OrderBy("Like", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, err
	}
	for _, i := range listDoc {
		var q response.House
		err = i.DataTo(&q)
		q.HouseID = i.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, nil
}

func (this *House) PutItem() (string, error) {
	//add to collection
	res, _, err := Client.Collection(this.GetCollectionKey()).Add(Ctx, *this)
	go searchIndex.SaveObject(HouseSearch{
		ObjectID:       res.ID,
		NearBy:         this.NearBy,
		Header:         this.Header,
		Price: 			this.Price,
	})
	if err != nil {
		return "", err
	}
	//add to search

	return res.ID, err
}

func (this *House) Public(time PostTime) error {
	//get item
	doc, err := this.GetCollection().Doc(time.HouseID).Get(Ctx)
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
	_, err = this.GetCollection().Doc(time.HouseID).Set(Ctx, h)
	return err
}

func (this *House) Delete(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Delete(Ctx)
	if err != nil {
		return err
	}
	_, err = searchIndex.Delete(id)
	return err
}

func (this *House) AddImage(file multipart.File) (string, error) {
	ref, _, err := Client.Collection(consts.IMAGE_LINK).Add(Ctx, map[string]string{
		"Image" : "link",
	})
	if err != nil {
		return "", err
	}
	wc := bucket.Object(ref.ID).NewWriter(Ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	return ref.ID, nil
}

func (this *House) GetFromKey(key string) (error) {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(Ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(this)
}

func (this *House) GetResponse(key string) (response.House, error) {
	doc, err := Client.Collection(this.GetCollectionKey()).Doc(key).Get(Ctx)
	if err != nil {
		return response.House{}, err
	}
	res := response.House{}
	err = doc.DataTo(&res)
	res.HouseID = doc.Ref.ID
	return res, err
}

func (this *House) GetAllActivate() ([]response.House, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Where("ExpiredTime", ">", time.Now().Unix()).Documents(Ctx)
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

func (this *House) GetPageActivate(page, count int) ([]response.House, int, error) {
	listdoc, err := Client.Collection(this.GetCollectionKey()).Where("ExpiredTime", ">", time.Now().Unix()).OrderBy("ExpiredTime", firestore.Asc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}
	total := len(listdoc)
	if page * count > total {
		return nil, 0, response.BadRequest
	}
	end := (page + 1) * count
	listHouse := []response.House{}
	if end > total {
		end = total
	}
	for _, i := range listdoc[page * count : end] {
		var q response.House
		err = i.DataTo(&q)
		q.HouseID = i.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, total, nil
}

func (this *House) GetAll() ([]response.House, error) {
	listdoc := Client.Collection(this.GetCollectionKey()).Documents(Ctx)
	listHouse := []response.House{}
	for {
		var q response.House
		doc, err := listdoc.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			continue
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
	listdoc := Client.Collection(this.GetCollectionKey()).Where("OwnerID", "==", id).Documents(Ctx)
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

func (this *House) GetPaginateHouseOfUser(id string, page int, count int) ([]response.House, int, error) {
	listHouse := []response.House{}
	start := page * count
	end := start + count

	list, err := this.GetCollection().Where("OwnerID", "==", id).OrderBy("PostTime", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}

	if start > len(list) {
		return nil, 0 , response.BadRequest
	}
	if end > len(list) {
		end = len(list)
	}
	total := len(list)
	for _, i := range list[start : end] {
		var q response.House
		err = i.DataTo(&q)
		q.HouseID = i.Ref.ID
		listHouse = append(listHouse, q)
	}
	return listHouse, total, nil
}

func (this *House) GetAllByStatus(status Status) ([]response.House, error) {
	listdoc := Client.Collection(consts.HOUSE).Where("Status", "==", status).Documents(Ctx)
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

func (this *House) GetPaginateByStatus(status Status, page int, count int) ([]response.House, int, error) {
	listHouse := []response.House{}
	start := page * count
	end := start + count

	l, err := this.GetCollection().Where("Status", "==", status).OrderBy("PostTime", firestore.Desc).Documents(Ctx).GetAll()
	if err != nil {
		return nil, 0, err
	}

	if start > len(l) {
		return nil, 0 , response.BadRequest
	}
	if end > len(l) {
		end = len(l)
	}
	for _, i := range l[start : end] {
		h := response.House{}
		err = i.DataTo(&h)
		if err != nil {
			continue
		}
		h.HouseID = i.Ref.ID
		listHouse = append(listHouse, h)
	}
	return listHouse, len(l), nil
}

func (this *House) UpdateItem(id string) error {
	_, err := Client.Collection(this.GetCollectionKey()).Doc(id).Set(Ctx, *this)
	if err != nil {
		return err
	}
	go searchIndex.SaveObject(HouseSearch{
		ObjectID:       id,
		NearBy:         this.NearBy,
		Header:         this.Header,
		Price: 			this.Price,
	})
	return nil
}

func (this *House) GetByPriceRange(startPrice, endPrice int) ([]response.House, error) {
	res := []response.House{}
	list, err := this.GetCollection().Where("Price", ">=", startPrice).Where("Price", "<=", endPrice).Documents(Ctx).GetAll()
	if err != nil {
		return res, err
	}
	for _, i := range list {
		h := response.House{}
		err = i.DataTo(&h)
		if err != nil {
			continue
		}
		h.HouseID = i.Ref.ID
		res = append(res, h)
	}
	return res, nil
}

func (this *House) SearchAllItem(key string, startPrice, endPrice int) ([]response.House, error) {
	log.Println(key, startPrice, endPrice, "   models/house_model.go:451")
	res, err := searchIndex.Search(key, opt.NumericFilter("price:" + strconv.Itoa(startPrice) + " TO " + strconv.Itoa(endPrice)))
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
		if resH.ExpiredTime > now {
			list = append(list, resH)
		}
	}
	return list, nil
}
