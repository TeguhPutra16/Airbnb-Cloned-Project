package reservation

import (
	"time"

	"gorm.io/gorm"
)

type CoreReservation struct {
	gorm.Model
	ID         uint
	HomestayID uint
	UserID     uint
	CheckIn    string
	CheckOut   string
	Price      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceEntities interface {
	GetAllhomestay() (data []CoreReservation, err error)
	Create(input CoreReservation) (err error)
	Update(id int, input CoreReservation) error
	GetById(id int) (data CoreReservation, err error)
	GethHomestaybyidUser(user_id int) (data []CoreReservation, err error)
	DeleteById(id int) (CoreReservation, error)
}

type RepositoryEntities interface {
	GetAllhomestay() (data []CoreReservation, err error)
	Create(input CoreReservation) (err error)
	Update(id int, input CoreReservation) error
	GetById(id int) (data CoreReservation, err error)
	GethHomestaybyidUser(user_id int) (data []CoreReservation, err error)
	DeleteById(id int) (CoreReservation, error)
}
