package repository

import (
	"be13/project/features/user"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryEntities { // user.repository mengimplementasikan interface repository yang ada di entities
	return &userRepository{
		db: db,
	}

}

// Create implements user.RepositoryEntities
func (repo *userRepository) Create(input user.CoreUser) (err error) {
	userGorm := FromUserCore(input) //dari gorm model ke user core yang ada di entities

	tx := repo.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil

}
func (repo *userRepository) GetAll() (data []user.CoreUser, err error) {
	var users []User //mengambil data gorm model(model.go)
	tx := repo.db.Unscoped().Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(users) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil

}
func (repo *userRepository) Update(id int, input user.CoreUser) error {
	var users User

	tx1 := repo.db.First(&users, id)

	if tx1.Error != nil {

		return tx1.Error
	}

	if input.Password == "" {
		input.Password = users.Password
	}

	userGorm := FromUserCore(input)
	input.Password = user.Bcript(input.Password)

	tx := repo.db.Model(&userGorm).Where("id = ?", id).Updates(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteById implements user.RepositoryEntities
func (repo *userRepository) DeleteById(id int) (user.CoreUser, error) {

	users := User{}
	users.Status = "Deactivated"
	tx := repo.db.Model(&users).Where("id = ?", id).Updates(&users)

	if tx.Error != nil {
		return user.CoreUser{}, tx.Error
	}

	tx1 := repo.db.Delete(&users, id)
	if tx1.Error != nil {
		return user.CoreUser{}, tx.Error
	}

	txres := repo.db.Unscoped().Where("id=?", id).Find(&users)
	if txres.Error != nil {
		return user.CoreUser{}, txres.Error
	}
	if tx.RowsAffected == 0 {
		return user.CoreUser{}, errors.New("id not found")

	}
	result := users.ModelsToCore()
	return result, nil
}

// // GetById implements user.RepositoryEntities
func (repo *userRepository) GetById(id int) (data user.CoreUser, err error) {
	var users User

	tx := repo.db.First(&users, id)

	if tx.Error != nil {

		return user.CoreUser{}, tx.Error
	}
	gorms := users.ModelsToCore()
	return gorms, nil
}
