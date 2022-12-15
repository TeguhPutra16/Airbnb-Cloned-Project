package repository

import (
	"be13/project/features/reservation"
	"be13/project/features/reservation/service"
	"log"

	"gorm.io/gorm"
)

type reservasiRepository struct {
	db *gorm.DB
}

func NewRes(db *gorm.DB) reservation.RepositoryEntities { // user.repository mengimplementasikan interface repository yang ada di entities
	return &reservasiRepository{
		db: db,
	}

}

// Create implements reservation.RepositoryEntities
func (repo *reservasiRepository) Create(input reservation.CoreReservation) (err error) {
	model := Homestay{}
	tx := repo.db.First(&model, input.HomestayID)
	if tx.Error != nil {
		return tx.Error
	}

	log.Println("harga", input.HomestayID)

	cost := service.WaktuNginap(input.CheckOut, input.CheckIn) * float64(model.Price)
	input.Price = cost
	log.Println("harga", input.Price)

	reservasi := FromCore(input) //dari gorm model ke user core yang ada di entities
	reservasi.Price = cost

	tx1 := repo.db.Create(&reservasi) // proses insert data

	if tx1.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteById implements reservation.RepositoryEntities
func (*reservasiRepository) DeleteById(id int) (reservation.CoreReservation, error) {
	panic("unimplemented")
}

// GetAllreservasi implements reservation.RepositoryEntities
func (repo *reservasiRepository) GetAllreservasi(user int) (data []reservation.CoreReservation, err error) {
	var res []Reservation //mengambil data gorm model(model.go)
	tx := repo.db.Where("user_id=?", user).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(res) //mengambil data dari gorm model(file repository(model.go))

	return DataCore, nil
}

// GetById implements reservation.RepositoryEntities
func (*reservasiRepository) GetById(id int) (data reservation.CoreReservation, err error) {
	panic("unimplemented")
}

// Update implements reservation.RepositoryEntities
func (*reservasiRepository) Update(id int, input reservation.CoreReservation) error {
	panic("unimplemented")
}
