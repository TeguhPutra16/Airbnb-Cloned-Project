package service

import (
	"be13/project/features/user"

	"errors"

	"github.com/go-playground/validator/v10"
)

// bisnis Logic
type UserService struct {
	userRepository user.RepositoryEntities //data repository dri entities
	validate       *validator.Validate
}

func New(repo user.RepositoryEntities) user.ServiceEntities { //dengan kembalian user.service
	return &UserService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceEntities
func (service *UserService) Create(input user.CoreUser) (err error) {
	input.Role = "User"
	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}

	return nil
}

// GetAll implements user.ServiceEntities
func (service *UserService) GetAll() (data []user.CoreUser, err error) {

	data, err = service.userRepository.GetAll() // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// Update implements user.ServiceEntities
func (service *UserService) Update(id int, input user.CoreUser) error {
	input.Password = user.Bcript(input.Password)
	errUpdate := service.userRepository.Update(id, input)
	if errUpdate != nil {
		return errors.New("GAGAL mengupdate data , QUERY ERROR")
	}

	return nil
}

// DeleteById implements user.ServiceEntities
func (service *UserService) DeleteById(id int) (user.CoreUser, error) {
	data, err := service.userRepository.DeleteById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return data, err
}

// // GetById implements user.ServiceEntities
func (service *UserService) GetById(id int) (data user.CoreUser, err error) {
	data, err = service.userRepository.GetById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}
