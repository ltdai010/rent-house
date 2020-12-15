package searchcontroller

import (
	"github.com/astaxie/beego"
	"rent-house/restapi/response"
	"rent-house/services/houseservices"
)

type SearchController struct {
	beego.Controller
}

// @Title GetAllSearchHouse
// @Description price range: VeryLow = "* - 500" DownLow = "500 - 700"  UpLow = "700 - 1000" DownMedium = "1000 - 1500" UpMedium = "1500 - 2000" High = "2000 - 2500" VeryHigh = "2500 - 3500" ExHigh = "3500 - 5000" More = "5000 - *""
// @Param	province	query	string	false	"province id"
// @Param	commune		query	string	false	"commune id"
// @Param	district	query	string	false	"district id"
// @Param	priceRange	query	string	false	"price range"
// @Param	key			query	string	true	"key"
// @Success 200 {object} models.House
// @Failure 403 :houseID is empty
// @router /search-results [get]
func (u *SearchController) GetAllSearchHouse() {
	key := u.GetString("key", "")
	provinceID := u.GetString("province", "")
	districtID := u.GetString("district", "")
	commune := u.GetString("commune", "")
	houses, err := houseservices.SearchHouse(key, provinceID, districtID, commune)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonSingle{
			Data: houses,
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}

// @Title GetPageActivateSearchHouse
// @Description price range: VeryLow = "* - 500" DownLow = "500 - 700" UpLow = "700 - 1000" DownMedium = "1000 - 1500" UpMedium = "1500 - 2000" High = "2000 - 2500" VeryHigh = "2500 - 3500" ExHigh = "3500 - 5000" More = "5000 - *""
// @Param	key			query	string	true		"key for search"
// @Param	province	query	string 	false		"province id"
// @Param	commune		query	string	false		"commune id"
// @Param	district	query	string	false		"distric id"
// @Param	priceRange	query	string	false		"price range"
// @Param	page		query	int		true		"page"
// @Param	count		query	int		true		"count"
// @Success 200 {object} models.House
// @router /page-search-results [get]
func (u *SearchController) GetPageActivateSearchHouse() {
	key := u.GetString("key")
	page, err := u.GetInt("page")
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	count, err := u.GetInt("count")
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
		u.ServeJSON()
		return
	}
	provinceID := u.GetString("province")
	districtID := u.GetString("district")
	commune := u.GetString("commune")
	users, total, err := houseservices.SearchPageHouse(key, provinceID, districtID, commune, page, count)
	if err != nil {
		u.Data["json"] = response.NewErr(response.BadRequest)
	} else {
		u.Data["json"] = response.ResponseCommonArray{
			Data: users,
			TotalCount: int64(total),
			Err:  response.NewErr(response.Success),
		}
	}
	u.ServeJSON()
}
