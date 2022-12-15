package repository

import (
	"be13/project/features/check"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type checkRepository struct {
	db *gorm.DB
}

func NewCheck(db *gorm.DB) check.RepositoryInterface {
	return &checkRepository{
		db: db,
	}
}

// GetAllhomestay implements check.RepositoryInterface
func (repo *checkRepository) Checkroom(id int, check_in string, check_out string) (string, error) {

	var check []Reservation
	checkIn, errConvtime1 := time.Parse("02/01/2006", check_in)
	if errConvtime1 != nil {
		return "Error Convert time 1", errConvtime1
	}
	log.Println("ini", check_in)
	checkOut, errConvtime2 := time.Parse("02/01/2006", check_out)
	if errConvtime2 != nil {
		return "Error Convert time 2", errConvtime2
	}

	tx := repo.db.Select("id").Where("homestay_id=? AND check_in BETWEEN ? AND ?", id, checkIn, checkOut).Or("homestay_id=? AND check_out BETWEEN ? AND ?", id, checkIn, checkOut).Find(&check) //start dan end harus di convert dulu
	if tx.Error != nil {
		return "", errors.New("error get homestay_id go:35")
	}

	if tx.RowsAffected == 0 {
		return "Available", nil
	}

	return "Not-Available", nil
}
