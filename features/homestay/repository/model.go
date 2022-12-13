package repository

import (
	"be13/project/features/homestay"

	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	Title       string
	Description string
	Address     string
	Status      string
	AvgRate     int
	Price       int
	UserID      uint
}

func FromCore(dataCore homestay.CoreHomestay) Homestay { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	homeGorm := Homestay{
		Title:       dataCore.Title,
		Description: dataCore.Description,
		Address:     dataCore.Address,
		Status:      dataCore.Status,
		AvgRate:     dataCore.AvgRate,
		Price:       dataCore.Price,
		UserID:      dataCore.ID,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return homeGorm //insert user
}
func (dataModel *Homestay) ModelsToCore() homestay.CoreHomestay { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return homestay.CoreHomestay{
		ID:          dataModel.ID,
		Title:       dataModel.Title,
		Description: dataModel.Description,
		Address:     dataModel.Address,
		Status:      dataModel.Status,
		AvgRate:     dataModel.AvgRate,
		Price:       dataModel.Price,
		UserID:      dataModel.ID,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Homestay) []homestay.CoreHomestay { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []homestay.CoreHomestay
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
