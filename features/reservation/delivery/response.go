package delivery

import (
	"be13/project/features/reservation"
	"time"
)

type ReservationRespon struct {
	ID         uint      `json:"id"`
	HomestayID uint      `json:"homestay_id"`
	UserID     uint      `json:"user_id"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func CoreToRespon(dataCore reservation.CoreReservation) ReservationRespon { // data user core yang ada di controller yang memanggil user repository
	return ReservationRespon{
		ID:         dataCore.ID,
		HomestayID: dataCore.HomestayID,
		UserID:     dataCore.UserID,
		CheckIn:    dataCore.CheckIn,
		CheckOut:   dataCore.CheckOut,
		Price:      dataCore.Price,
		CreatedAt:  dataCore.CreatedAt,
		UpdatedAt:  dataCore.UpdatedAt,
	}
}
func ListCoreToRespon(dataCore []reservation.CoreReservation) []ReservationRespon { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []ReservationRespon

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, CoreToRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
