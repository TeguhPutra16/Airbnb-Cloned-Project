package repository

import (
	"be13/project/features/images"

	"gorm.io/gorm"
)

type Upload struct {
	gorm.Model
	Images     string `gorm:"type:varchar(255);not null"`
	FileName   string
	HomestayID uint
}

func FromCore(dataCore images.CoreUpload) Upload { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	uploadGorm := Upload{
		FileName:   dataCore.FileName,
		Images:     dataCore.Images, //mapping data core ke data gorm model
		HomestayID: dataCore.HomestayID,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return uploadGorm //insert user
}
func (dataModel *Upload) ModelsToCore() images.CoreUpload { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return images.CoreUpload{
		ID:         dataModel.ID,
		FileName:   dataModel.FileName,
		Images:     dataModel.Images, //mapping data core ke data gorm model
		HomestayID: dataModel.HomestayID,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}
func ListModelTOCore(dataModel []Upload) []images.CoreUpload { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []images.CoreUpload
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
