package comment

import "time"

type CoreComment struct {
	ID         uint      `json:"id"`
	HomestayID uint      `json:"homestay_id"`
	UserID     uint      `json:"user_id"`
	Notes      string    `json:"notes"`
	Ratings    int       `json:"ratings"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ServiceInterface interface {
	GetAllComment() (data []CoreComment, err error)
	CreateComment(input CoreComment) (err error)
	UpdateComment(id int, userid int, input CoreComment) error
	GetById(id int) (data CoreComment, err error)
	DeleteById(id int, userid int) (CoreComment, error)
}

type RepositoryInterface interface {
	GetAllComment() (data []CoreComment, err error)
	CreateComment(input CoreComment) (err error)
	UpdateComment(id int, userid int, input CoreComment) error
	GetById(id int) (data CoreComment, err error)
	DeleteById(id int, userid int) (CoreComment, error)
}
