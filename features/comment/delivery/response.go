package delivery

import (
	"be13/project/features/comment"
	"time"
)

type CommentRespon struct {
	ID         uint
	HomestayID uint
	UserID     uint
	Notes      string
	Ratings    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func CoreToRespon(dataCore comment.CoreComment) CommentRespon { // data user core yang ada di controller yang memanggil user repository
	return CommentRespon{
		ID:         dataCore.ID,
		HomestayID: dataCore.HomestayID,
		UserID:     dataCore.UserID,
		Notes:      dataCore.Notes,
		Ratings:    dataCore.Ratings,
		CreatedAt:  dataCore.CreatedAt,
		UpdatedAt:  dataCore.UpdatedAt,
	}
}
func ListCoreToRespon(dataCore []comment.CoreComment) []CommentRespon { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []CommentRespon

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, CoreToRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
