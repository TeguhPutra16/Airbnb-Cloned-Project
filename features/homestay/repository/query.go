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
func (repo *homeStayRepository) Create(input homestay.CoreHomestay) error {

	home := FromCore(input) //dari gorm model ke user core yang ada di entities

	tx := repo.db.Create(&home) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	user := User{}
	user.Status = "Hosting"

	tx1 := repo.db.Model(&user).Where("id = ?", input.UserID).Updates(&user)
	if tx1.RowsAffected == 0 {
		return errors.New("error update role")
	}

	return nil
}

// DeleteById implements homestay.RepositoryEntities
func (repo *homeStayRepository) DeleteById(id int, userId int) (homestay.CoreHomestay, error) {
	homes := Homestay{}
	tx := repo.db.Where("id=?", id).Find(&homes)
	if tx.Error != nil {
		return homestay.CoreHomestay{}, tx.Error
	}

	if homes.UserID != uint(userId) {
		return homestay.CoreHomestay{}, errors.New("hanya bisa ubah data sendiri")
	}
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
func (repo *homeStayRepository) GetAll() (data []homestay.CoreHomestay, err error) {
	var home []Homestay //mengambil data gorm model(model.go)
	tx := repo.db.Find(&home)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(home) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil
}

// GetById implements homestay.RepositoryEntities
func (repo *homeStayRepository) GetById(id int) (data homestay.CoreHomestay, err error) {
	var home Homestay

	tx := repo.db.Preload("Comments").First(&home, id)

	if tx.Error != nil {

		return homestay.CoreHomestay{}, tx.Error
	}
	gorms := home.ModelsToCore()
	return gorms, nil
}

// GetBytime implements homestay.RepositoryEntities

// Update implements homestay.RepositoryEntities
func (repo *homeStayRepository) Update(id int, userId int, input homestay.CoreHomestay) error {
	var homes Homestay
	home := FromCore(input)
	tx1 := repo.db.Where("id=?", id).Find(&homes)
	tx := repo.db.Model(&home).Where("id = ?", id).Updates(&home)

	if homes.UserID != uint(userId) {
		return errors.New("hanya bisa ubah data sendiri")
	}

	if tx.Error != nil {
		return tx.Error
	}
	if tx1.Error != nil {
		return tx1.Error
	}
	return nil
}

// GethHomestaybyidUser implements homestay.RepositoryEntities
func (repo *homeStayRepository) GethHomestaybyidUser(user_id int) (data []homestay.CoreHomestay, err error) {
	var home []Homestay //mengambil data gorm model(model.go)
	tx := repo.db.Where("user_id=?", user_id).Find(&home)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(home) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil
}
