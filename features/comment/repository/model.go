package repository

import (
	"be13/project/features/comment"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	HomestayID uint
	UserID     uint
	Notes      string
	Ratings    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func FromCore(dataCore comment.CoreComment) Comment { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	homeGorm := Comment{
		HomestayID: dataCore.HomestayID,
		UserID:     dataCore.UserID,
		Notes:      dataCore.Notes,
		Ratings:    dataCore.Ratings,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return homeGorm //insert user
}
func (dataModel *Comment) ModelsToCore() comment.CoreComment { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return comment.CoreComment{
		ID:         dataModel.ID,
		HomestayID: dataModel.HomestayID,
		UserID:     dataModel.UserID,
		Notes:      dataModel.Notes,
		Ratings:    dataModel.Ratings,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Comment) []comment.CoreComment { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []comment.CoreComment
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
