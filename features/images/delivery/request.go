package delivery

import (
	"be13/project/features/images"
)

type UploadRequest struct {
	Images     string
	FileName   string
	HomestayID uint
}

func RequestToCore(data UploadRequest) images.CoreUpload {
	return images.CoreUpload{
		FileName:   data.FileName,
		Images:     data.Images, //mapping data core ke data gorm model
		HomestayID: data.HomestayID,
	}
}
