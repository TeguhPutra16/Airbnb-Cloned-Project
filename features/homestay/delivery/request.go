package delivery

import (
	"be13/project/features/homestay"
)

type HomestayRequest struct {
	Address     string `json:"address" form:"address"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Images      string `json:"images" form:"images"`
}

func UserRequestToUserCore(data HomestayRequest) homestay.CoreHomestay {
	return homestay.CoreHomestay{
		Title:       data.Title,
		Description: data.Description,
		Address:     data.Address,
		Price:       data.Price,
		Images:      data.Images,
	}
}
