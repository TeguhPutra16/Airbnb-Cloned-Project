package images

import "time"

type CoreUpload struct {
	ID         uint
	Images     string `gorm:"type:varchar(255);not null"`
	FileName   string
	HomestayID uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceEntities interface { //sebagai contract yang dibuat di layer service
	GetAll() (data []CoreUpload, err error) //yang returnnya(mengembalikan data core)
	UplodaImg(input CoreUpload) (err error) // menambahkah data user berdasarkan data usercore

	Update(id int, input CoreUpload) error

	GetById(id int) (data CoreUpload, err error)
	DeleteById(id int) (CoreUpload, error)
}

type RepositoryEntities interface { // berkaitan database
	GetAll() (data []CoreUpload, err error) //yang returnnya(mengembalikan data core)
	UplodaImg(input CoreUpload) (err error) // menambahkah data user berdasarkan data usercore

	Update(id int, input CoreUpload) error

	GetById(id int) (data CoreUpload, err error)
	DeleteById(id int) (CoreUpload, error)
}
