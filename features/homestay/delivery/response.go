package delivery

import (
	"be13/project/features/homestay"
	"time"
)

type HomestayRespon struct {
	ID          uint
	Title       string
	Description string
	Address     string
	Status      string
	AvgRate     uint
	Price       uint
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CoreToRespon(dataCore homestay.CoreHomestay) HomestayRespon { // data user core yang ada di controller yang memanggil user repository
	return HomestayRespon{
		ID:          dataCore.ID,
		Title:       dataCore.Title,
		Description: dataCore.Description,
		Address:     dataCore.Address,
		Status:      dataCore.Status,
		AvgRate:     dataCore.AvgRate,
		Price:       dataCore.Price,
		UserID:      dataCore.ID,
		CreatedAt:   dataCore.CreatedAt,
		UpdatedAt:   dataCore.UpdatedAt,
	}
}
func ListCoreToRespon(dataCore []homestay.CoreHomestay) []HomestayRespon { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []HomestayRespon

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, CoreToRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}