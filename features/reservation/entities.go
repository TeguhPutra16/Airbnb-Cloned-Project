package reservation

import (
	"time"
)

type CoreReservation struct {
	ID         uint
	HomestayID uint
	UserID     uint
	CheckIn    time.Time
	CheckOut   time.Time
	Price      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceEntities interface {
	GetAllreservasi(user int) (data []CoreReservation, err error)
	Create(input CoreReservation) (err error)
	Update(id int, input CoreReservation) error
	GetById(id int) (data CoreReservation, err error)
	DeleteById(id int) (CoreReservation, error)
}

type RepositoryEntities interface {
	GetAllreservasi(user int) (data []CoreReservation, err error)
	Create(input CoreReservation) (err error)
	Update(id int, input CoreReservation) error
	GetById(id int) (data CoreReservation, err error)
	DeleteById(id int) (CoreReservation, error)
}
