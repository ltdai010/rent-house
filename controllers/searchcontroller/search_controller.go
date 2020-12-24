package searchcontroller

import (
	"errors"
	"github.com/astaxie/beego"
	"rent-house/restapi/response"
	"rent-house/services/houseservices"
)

type SearchController struct {
	beego.Controller
}


// @Title GetPageActivateSearchHouse
// @Description price range: VeryLow = "* - 500" DownLow = "500 - 700" UpLow = "700 - 1000" DownMedium = "1000 - 1500" UpMedium = "1500 - 2000" High = "2000 - 2500" VeryHigh = "2500 - 3500" ExHigh = "3500 - 5000" More = "5000 - *""
// @Param	key			query	string	true		"key for search"
// @Param	province	query	string 	false		"province id"
// @Param	commune		query	string	false		"commune id"
// @Param	district	query	string	false		"district id"
// @Param	price		query	string	false		"price range"
// @Param	page		query	int		true		"page"
// @Param	count		query	int		true		"count"
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
	users := []response.House{}
	total := 0
	err := errors.New("")
	if count <= 0 {
		users, err = houseservices.SearchHouse(key, provinceID, districtID, commune, price)
		total = len(users)
	} else {
		users, total, err = houseservices.SearchPageHouse(key, provinceID, districtID, commune, price, page, count)

	}
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonArray{
			Data:       users,
			TotalCount: int64(total),
			Err:        response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}
