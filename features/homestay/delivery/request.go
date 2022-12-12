package delivery

import (
	"be13/project/features/homestay"
)

type HomestayRequest struct {
	Address     string `json:"address" form:"address"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	AvgRate     uint   `json:"avgRate" form:"avgRate"`
	Price       uint   `json:"price" form:"price"`
}

func UserRequestToUserCore(data HomestayRequest) homestay.CoreHomestay {
	return homestay.CoreHomestay{
		Title:       data.Title,
		Description: data.Description,
		Address:     data.Address,
		AvgRate:     data.AvgRate,
		Price:       data.Price,
	}
}
