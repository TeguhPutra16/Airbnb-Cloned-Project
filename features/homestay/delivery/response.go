package delivery

import (
	"be13/project/features/homestay"
	"time"
)

type HomestayRespon struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	AvgRate     int       `json:"avg_rate"`
	Price       int       `json:"price"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CoreToRespon(dataCore homestay.CoreHomestay) HomestayRespon { // data user core yang ada di controller yang memanggil user repository
	return HomestayRespon{
		ID:          dataCore.ID,
		Title:       dataCore.Title,
		Description: dataCore.Description,
		Address:     dataCore.Address,
		AvgRate:     dataCore.AvgRate,
		Price:       dataCore.Price,
		UserID:      dataCore.UserID,
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
