package repository

import (
	"be13/project/features/check"
	"be13/project/features/homestay"

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
func (*checkRepository) GetAllhomestay(check_in string, checkout_out string) (data []homestay.CoreHomestay, err error) {
	panic("unimplemented")
}
