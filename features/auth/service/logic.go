package service

import (
	"be13/project/features/auth"
	"be13/project/features/user/repository"
)

type authService struct {
	authRepository auth.RepositoryInterface
}

// Login implements auth.ServiceInterface

func NewAuth(repo auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authRepository: repo,
	}
}

func (service *authService) Login(email string, pass string) (string, repository.User, error) {
	token, data, err := service.authRepository.Login(email, pass)
	return token, data, err
}
