package delivery

import (
	"be13/project/features/images"
	"time"
)

type UploadResponse struct {
	ID         uint
	Images     string `gorm:"type:varchar(255);not null"`
	FileName   string
	HomestayID uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func CoreToRespon(dataCore images.CoreUpload) UploadResponse { // data user core yang ada di controller yang memanggil user repository
	return UploadResponse{
		ID:         dataCore.ID,
		FileName:   dataCore.FileName,
		Images:     dataCore.Images, //mapping data core ke data gorm model
		HomestayID: dataCore.HomestayID,
		CreatedAt:  dataCore.CreatedAt,
		UpdatedAt:  dataCore.UpdatedAt,
	}
}
func ListCoreToRespon(dataCore []images.CoreUpload) []UploadResponse { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []UploadResponse

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, CoreToRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
