package convert

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func ConvertRenterResponse(id string, renter models.Renter) response.Renter  {
	return response.Renter{
		RenterID:      id,
		RenterName:    renter.RenterName,
		PhoneNumber:   renter.PhoneNumber,
		Email:         renter.Email,
		ListFavourite: renter.ListFavourite,
	}
}
