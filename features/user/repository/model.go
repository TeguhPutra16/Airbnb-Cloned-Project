package repository

import (
	comment "be13/project/features/comment/repository"
	"be13/project/features/homestay/repository"
	_user "be13/project/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50)"`
	Password  string `gorm:"type:varchar(500)"`
	Email     string
	Address   string
	Status    string
	Role      string
	Homestays []repository.Homestay
	Comments  []comment.Comment
}

func FromUserCore(dataCore _user.CoreUser) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email, //mapping data core ke data gorm model
		Password: dataCore.Password,
		Address:  dataCore.Address,
		Role:     dataCore.Role,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return userGorm //insert user
}
func (dataModel *User) ModelsToCore() _user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return _user.CoreUser{
		ID:       dataModel.ID,
		Name:     dataModel.Name,
		Email:    dataModel.Email, //mapping data core ke data gorm model
		Password: dataModel.Password,
		Address:  dataModel.Address,
		Role:     dataModel.Role,
	}
}
func ListModelTOCore(dataModel []User) []_user.CoreUser { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []_user.CoreUser
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
