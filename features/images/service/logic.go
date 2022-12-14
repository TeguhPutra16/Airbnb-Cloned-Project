package service

import (
	"be13/project/features/images"

	"github.com/go-playground/validator/v10"
)

// bisnis Logic
type imageService struct {
	imageRepository images.RepositoryEntities //data repository dri entities
	validate        *validator.Validate
}

func NewImage(repo images.RepositoryEntities) images.ServiceEntities { //dengan kembalian user.service
	return &imageService{
		imageRepository: repo,
		validate:        validator.New(),
	}
}

// DeleteById implements images.ServiceEntities
func (*imageService) DeleteById(id int) (images.CoreUpload, error) {
	panic("unimplemented")
}

// GetAll implements images.ServiceEntities
func (*imageService) GetAll() (data []images.CoreUpload, err error) {
	panic("unimplemented")
}

// GetById implements images.ServiceEntities
func (*imageService) GetById(id int) (data images.CoreUpload, err error) {
	panic("unimplemented")
}

// Update implements images.ServiceEntities
func (*imageService) Update(id int, input images.CoreUpload) error {
	panic("unimplemented")
}

// UplodaImg implements images.ServiceEntities
func (service *imageService) UplodaImg(input images.CoreUpload) (err error) {
	err = service.imageRepository.UplodaImg(input)
	if err != nil {
		return err
	}
	return nil
}
