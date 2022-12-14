package repository

import (
	"be13/project/features/images"

	"gorm.io/gorm"
)

type imageRepository struct {
	db *gorm.DB
}

func NewImage(db *gorm.DB) images.RepositoryEntities { // user.repository mengimplementasikan interface repository yang ada di entities
	return &imageRepository{
		db: db,
	}

}

// DeleteById implements images.RepositoryEntities
func (*imageRepository) DeleteById(id int) (images.CoreUpload, error) {
	panic("unimplemented")
}

// GetAll implements images.RepositoryEntities
func (*imageRepository) GetAll() (data []images.CoreUpload, err error) {
	panic("unimplemented")
}

// GetById implements images.RepositoryEntities
func (*imageRepository) GetById(id int) (data images.CoreUpload, err error) {
	panic("unimplemented")
}

// Update implements images.RepositoryEntities
func (*imageRepository) Update(id int, input images.CoreUpload) error {
	panic("unimplemented")
}

// UplodaImg implements images.RepositoryEntities
func (repo *imageRepository) UplodaImg(input images.CoreUpload) (err error) {
	dataModel := FromCore(input)

	tx := repo.db.Create(&dataModel).Last(&dataModel)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
