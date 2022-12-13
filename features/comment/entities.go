package comment

import "time"

type CoreComment struct {
	ID         uint
	HomestayID uint `validate:"required"`
	UserID     uint
	Notes      string
	Ratings    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceInterface interface {
	GetAllComment() (data []CoreComment, err error)
	CreateComment(input CoreComment) (err error)
	UpdateComment(id int, input CoreComment) error
	GetById(id int) (data CoreComment, err error)
	DeleteById(id int) (CoreComment, error)
}

type RepositoryInterface interface {
	GetAllComment() (data []CoreComment, err error)
	CreateComment(input CoreComment) (err error)
	UpdateComment(id int, input CoreComment) error
	GetById(id int) (data CoreComment, err error)
	DeleteById(id int) (CoreComment, error)
}
