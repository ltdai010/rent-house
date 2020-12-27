package searchcontroller

import (
	"github.com/astaxie/beego"
	"rent-house/restapi/response"
	"rent-house/services/houseservices"
)

type SearchController struct {
	beego.Controller
}


// @Title GetPageActivateSearchHouse
// @Description price range: "* - 1000" "1000 - 2000" "2000 - 3500" "3500 - 5000" "5000 - *"
// @Param	key			query	string	false		"key for search"
// @Param	province	query	string 	false		"province id"
// @Param	commune		query	string	false		"commune id"
// @Param	district	query	string	false		"district id"
// @Param	price		query	string	false		"price range"
// @Param   house_type	query	int		false		"house type: 0:single room, 1:mini apartment, 2:full house, 3:apartment"
// @Param	page		query	int		false		"page"
// @Param	count		query	int		false		"count"
// @Success 200 {object} models.House
// @router /page-search-results [get]
func (u *SearchController) GetPageActivateSearchHouse() {
	key := u.GetString("key")
	page, _ := u.GetInt("page")
	count, _ := u.GetInt("count")
	provinceID := u.GetString("province")
	districtID := u.GetString("district")
	commune := u.GetString("commune")
	price := u.GetString("price")
	houseType, err := u.GetInt("house_type")
	if err != nil {
		houseType = -1
	}
	users := []response.House{}
	total := 0
	if count <= 0 {
		users, err = houseservices.SearchHouse(key, provinceID, districtID, commune, price, houseType)
		total = len(users)
	} else {
		users, total, err = houseservices.SearchPageHouse(key, provinceID, districtID, commune, price, page, count, houseType)
	}
	if err != nil {
		u.Data["json"] = response.NewErr(response.NotMore)
	} else {
		u.Data["json"] = response.ResponseCommonArray{
			Data:       users,
			TotalCount: int64(total),
			Err:        response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}
