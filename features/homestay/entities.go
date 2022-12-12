package homestay

import "time"

type CoreHomestay struct {
	ID          uint
	Title       string
	Description string
	Address     string
	Status      string
	AvgRate     uint
	Price       uint
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ServiceEntities interface {
	GetAllhomestay() (data []CoreHomestay, err error)
	Create(input CoreHomestay) (err error)
	Update(id int, input CoreHomestay) error
	GetById(id int) (data CoreHomestay, err error)
	GetBytime(start, end string) (data []CoreHomestay, err error)
	DeleteById(id int) (CoreHomestay, error)
}

type RepositoryEntities interface {
	GetAll() (data []CoreHomestay, err error)
	Create(input CoreHomestay) (err error)
	Update(id int, input CoreHomestay) error
	GetById(id int) (data CoreHomestay, err error)
	GetBytime(start, end string) (data []CoreHomestay, err error)
	DeleteById(id int) (CoreHomestay, error)
}
