package auth

import "be13/project/features/user/repository"

type ServiceInterface interface {
	Login(email string, pass string) (string, repository.User, error)
}

type RepositoryInterface interface {
	Login(email string, pass string) (string, repository.User, error)
}
