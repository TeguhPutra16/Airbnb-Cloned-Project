package repository

import (
	"be13/project/features/auth"
	"be13/project/features/user/repository"
	"be13/project/middlewares"
	"errors"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) auth.RepositoryInterface {
	return &authRepository{
		db: db,
	}
}

// Login implements auth.RepositoryInterface
func (repo *authRepository) Login(email string, pass string) (string, repository.User, error) {
	var userData repository.User
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return "", repository.User{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", repository.User{}, errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID), userData.Role)
	if errToken != nil {
		return "", repository.User{}, errToken
	}

	return token, userData, nil
}
