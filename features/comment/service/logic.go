package service

import (
	"be13/project/features/comment"

	"github.com/go-playground/validator/v10"
)

func NewComment(repo comment.RepositoryInterface) comment.ServiceInterface { //dengan kembalian user.service
	return &commentService{
		commentRepository: repo,
		validate:          validator.New(),
	}
}

// bisnis Logic
type commentService struct {
	commentRepository comment.RepositoryInterface //data repository dri entities
	validate          *validator.Validate
}

// CreateComment implements comment.ServiceInterface
func (service *commentService) CreateComment(input comment.CoreComment) (err error) {
	err = service.commentRepository.CreateComment(input)
	return err
}

// DeleteById implements comment.ServiceInterface
func (*commentService) DeleteById(id int) (comment.CoreComment, error) {
	panic("unimplemented")
}

// GetAllComment implements comment.ServiceInterface
func (*commentService) GetAllComment() (data []comment.CoreComment, err error) {
	panic("unimplemented")
}

// GetById implements comment.ServiceInterface
func (*commentService) GetById(id int) (data comment.CoreComment, err error) {
	panic("unimplemented")
}

// UpdateComment implements comment.ServiceInterface
func (*commentService) UpdateComment(id int, input comment.CoreComment) error {
	panic("unimplemented")
}
