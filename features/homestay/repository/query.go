package repository

import (
	"be13/project/features/homestay"
	"errors"
	"time"

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
func (repo *homeStayRepository) DeleteById(id int) (homestay.CoreHomestay, error) {
	home := Homestay{}
	tx1 := repo.db.Delete(&home, id)
	if tx1.Error != nil {
		return homestay.CoreHomestay{}, tx1.Error
	}

	txres := repo.db.Unscoped().Where("id=?", id).Find(&home)
	if txres.Error != nil {
		return homestay.CoreHomestay{}, txres.Error
	}
	if tx1.RowsAffected == 0 {
		return homestay.CoreHomestay{}, errors.New("id not found")

	}
	result := home.ModelsToCore()
	return result, nil
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
func (repo *homeStayRepository) GetBytime(start string, end string) (data []homestay.CoreHomestay, err error) {
	var home []Homestay
	checkIn, errConvtime1 := time.Parse("02/01/2006", start)
	if errConvtime1 != nil {
		return nil, errConvtime1
	}
	checkOut, errConvtime2 := time.Parse("02/01/2006", end)
	if errConvtime2 != nil {
		return nil, errConvtime2
	}

	tx := repo.db.Where("created_at BETWEEN ? AND ?AND status=?", checkIn, checkOut, "Available").Find(&home) //start dan end harus di convert dulu
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("login failed")
	}

	var DataCore = ListModelTOCore(home) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil
}

// Update implements homestay.RepositoryEntities
func (*homeStayRepository) Update(id int, input homestay.CoreHomestay) error {
	panic("unimplemented")
}
