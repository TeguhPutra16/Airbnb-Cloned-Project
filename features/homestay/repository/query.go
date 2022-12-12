package repository

import (
	"be13/project/features/homestay"
	"errors"

	"gorm.io/gorm"
)

type homeStayRepository struct {
	db *gorm.DB
}

func NewHome(db *gorm.DB) homestay.RepositoryEntities { // user.repository mengimplementasikan interface repository yang ada di entities
	return &homeStayRepository{
		db: db,
	}

}

// Create implements homestay.RepositoryEntities
func (repo *homeStayRepository) Create(input homestay.CoreHomestay) (err error) {

	userGorm := FromCore(input) //dari gorm model ke user core yang ada di entities

	tx := repo.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// DeleteById implements homestay.RepositoryEntities
func (*homeStayRepository) DeleteById(id int) (homestay.CoreHomestay, error) {
	panic("unimplemented")
}

// GetAll implements homestay.RepositoryEntities
func (*homeStayRepository) GetAll() (data []homestay.CoreHomestay, err error) {
	panic("unimplemented")
}

// GetById implements homestay.RepositoryEntities
func (*homeStayRepository) GetById(id int) (data homestay.CoreHomestay, err error) {
	panic("unimplemented")
}

// GetBytime implements homestay.RepositoryEntities
func (*homeStayRepository) GetBytime(start string, end string) (data []homestay.CoreHomestay, err error) {
	panic("unimplemented")
}

// Update implements homestay.RepositoryEntities
func (*homeStayRepository) Update(id int, input homestay.CoreHomestay) error {
	panic("unimplemented")
}
