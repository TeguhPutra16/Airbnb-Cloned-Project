package repository

import (
	"be13/project/features/reservation"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	HomestayID uint
	UserID     uint
	CheckIn    string
	CheckOut   string
	Price      int
}

func FromCore(dataCore reservation.CoreReservation) Reservation { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	reserveGorm := Reservation{
		HomestayID: dataCore.HomestayID,
		UserID:     dataCore.UserID,
		CheckIn:    dataCore.CheckIn,
		CheckOut:   dataCore.CheckOut,
		Price:      dataCore.Price,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return reserveGorm //insert user
}

func (dataModel *Reservation) ModelsToCore() reservation.CoreReservation { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return reservation.CoreReservation{
		ID:         dataModel.ID,
		HomestayID: dataModel.HomestayID,
		UserID:     dataModel.UserID,
		CheckIn:    dataModel.CheckIn,
		CheckOut:   dataModel.CheckOut,
		Price:      dataModel.Price,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Reservation) []reservation.CoreReservation { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []reservation.CoreReservation
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
