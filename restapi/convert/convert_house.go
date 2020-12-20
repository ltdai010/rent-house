package convert

import (
	"rent-house/models"
	"rent-house/restapi/response"
)

func ConvertHouseReponse(id string, house models.House) response.House {
	return response.House{
		HouseID:        id,
		OwnerID:        house.OwnerID,
		HouseType:      response.HouseType(house.HouseType),
		Price:          house.Price,
		Unit:           response.Unit(house.Unit),
		Address:        response.Address{
			Province: house.Address.Province,
			District: house.Address.District,
			Commune:  house.Address.Commune,
			Street:   house.Address.Street,
		},
		Infrastructure: response.Infrastructure{
			PrivateBathroom: house.Infrastructure.PrivateBathroom,
			Heater:          house.Infrastructure.Heater,
			AirCondition:    house.Infrastructure.AirCondition,
			Balcony:         house.Infrastructure.Balcony,
			ElectricPrice:   house.Infrastructure.ElectricPrice,
			WaterPrice:      house.Infrastructure.WaterPrice,
			NumberOfRoom:    house.Infrastructure.NumberOfRoom,
			Kitchen:         house.Infrastructure.Kitchen,
			Other:           house.Infrastructure.Other,
		},
		NearBy:         house.NearBy,
		PreOrder:       house.PreOrder,
		Surface:        house.Surface,
		WithOwner:      house.WithOwner,
		ImageLink:      house.ImageLink,
		LastViewed:     house.LastViewed,
		MonthlyView:    house.MonthlyView,
		Header:         house.Header,
		View:           house.View,
		Like:           house.Like,
		Rented:         house.Rented,
		Content:        house.Content,
		PostTime:       house.PostTime,
		Status:         response.Status(house.Status),
		Review:         house.Review,
		AppearTime:     house.AppearTime,
		ExpiredTime:    house.ExpiredTime,
		AdminComment:   house.AdminComment,
	}
}
