package homestay

import (
	"be13/project/features/comment"
	"time"
)

type CoreHomestay struct {
	ID          uint
	Title       string
	Description string
	Address     string
	AvgRate     float64
	Price       int
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Images      string
	Comments    []comment.CoreComment
}

type ServiceEntities interface {
	GetAllhomestay() (data []CoreHomestay, err error)
	Create(input CoreHomestay) (err error)
	Update(id int, userId int, input CoreHomestay) error
	GetById(id int) (data CoreHomestay, err error)
	GethHomestaybyidUser(user_id int) (data []CoreHomestay, err error)
	DeleteById(id int, userId int) (CoreHomestay, error)
}

type RepositoryEntities interface {
	GetAll() (data []CoreHomestay, err error)
	Create(input CoreHomestay) (err error)
	Update(id int, userId int, input CoreHomestay) error
	GetById(id int) (data CoreHomestay, err error)
	GethHomestaybyidUser(user_id int) (data []CoreHomestay, err error)
	DeleteById(id int, userId int) (CoreHomestay, error)
}
