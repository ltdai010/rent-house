package addressservices

import "rent-house/models"

func GetAllProvince() ([]models.Province, error) {
	p := &models.Province{}
	return p.GetAll()
}

func GetAllDistrict(provinceID string) ([]models.District, error) {
	d := &models.District{}
	return d.GetAll(provinceID)
}

func GetAllCommune(districtID string) ([]models.Commune, error) {
	c := &models.Commune{}
	return c.GetAll(districtID)
}
