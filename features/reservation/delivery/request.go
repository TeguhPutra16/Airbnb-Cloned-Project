package delivery

import (
	"be13/project/features/reservation"
	"time"
)

type ReservasiRequest struct {
	HomestayID uint    `json:"homestay_id" form:"homestay_id"`
	UserID     uint    `json:"user_id" form:"user_id"`
	CheckIn    string  `json:"check_in" form:"check_in"`
	CheckOut   string  `json:"check_out" form:"check_out"`
	Price      float64 `json:"price" form:"price"`
}

func RequestToCore(data ReservasiRequest, checkIn, checkOut time.Time) reservation.CoreReservation {
	return reservation.CoreReservation{
		HomestayID: data.HomestayID,
		UserID:     data.UserID,
		CheckIn:    checkIn,
		CheckOut:   checkOut,
		Price:      data.Price,
	}
}
