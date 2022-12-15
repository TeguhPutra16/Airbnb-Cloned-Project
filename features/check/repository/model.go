package repository

import (
	"be13/project/features/homestay"
	"time"

	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	Title       string
	Description string
	Address     string
	AvgRate     float64
	Price       int
	UserID      uint
	Images      string
	Comments    []Comment
	Checks      []Reservation
}
type Reservation struct {
	gorm.Model
	HomestayID uint
	UserID     uint
	CheckIn    time.Time
	CheckOut   time.Time
	Price      float64
}
type Comment struct {
	gorm.Model
	HomestayID uint
	UserID     uint
	Notes      string
	Ratings    int
}

func FromCore(dataCore homestay.CoreHomestay) Homestay { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	homeGorm := Homestay{
		Title:       dataCore.Title,
		Description: dataCore.Description,
		Address:     dataCore.Address,
		AvgRate:     dataCore.AvgRate,
		Price:       dataCore.Price,
		UserID:      dataCore.UserID,
		Images:      dataCore.Images,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return homeGorm //insert user
}

func (dataModel *Homestay) ModelsToCore() homestay.CoreHomestay { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return homestay.CoreHomestay{
		ID:          dataModel.ID,
		Title:       dataModel.Title,
		Description: dataModel.Description,
		Address:     dataModel.Address,
		AvgRate:     dataModel.AvgRate,
		Price:       dataModel.Price,
		UserID:      dataModel.UserID,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
		Images:      dataModel.Images,
	}
}
func ListModelTOCore(dataModel []Homestay) []homestay.CoreHomestay { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []homestay.CoreHomestay
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}

// func LoadFeedsModeltoCore(model []Comment) []comment.CoreComment {
// 	var core []comment.CoreComment
// 	for _, v := range model {
// 		core = append(core, v.ModelsToCore())
// 	}
// 	return core

// }

// func (dataModel *Comment) ModelsToCore() comment.CoreComment { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
// 	return comment.CoreComment{
// 		ID:         dataModel.ID,
// 		HomestayID: dataModel.HomestayID,
// 		UserID:     dataModel.UserID,
// 		Notes:      dataModel.Notes,
// 		Ratings:    dataModel.Ratings,
// 		CreatedAt:  dataModel.CreatedAt,
// 		UpdatedAt:  dataModel.UpdatedAt,
// 	}
// }

// type User struct {
// 	gorm.Model
// 	Name      string `gorm:"type:varchar(50)"`
// 	Password  string `gorm:"type:varchar(500)"`
// 	Email     string
// 	Address   string
// 	Status    string
// 	Role      string
// 	Homestays []Homestay
// 	Comments  []Comment
// }
