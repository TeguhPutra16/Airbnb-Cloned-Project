package repository

import (
	"be13/project/features/comment"
	"be13/project/features/homestay/repository"

	"log"

	"errors"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) comment.RepositoryInterface { // user.repository mengimplementasikan interface repository yang ada di entities
	return &commentRepository{
		db: db,
	}
}

// CreateComment implements comment.RepositoryInterface
func (repo *commentRepository) CreateComment(input comment.CoreComment) (err error) {

	commentGorm := FromCore(input)     //dari gorm model ke user core yang ada di entities
	tx := repo.db.Create(&commentGorm) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	/////////////////////////////////////////menghitung jumlah comment untuk suatu homestay//////////////////////////////////////////////////
	var commentModel []Comment

	repo.db.Where("homestay_id=?", input.HomestayID).Find(&commentModel)

	jumData := len(commentModel)
	log.Println("jumData", jumData)
	////////////////////////////////////////select curretn avgrating////////////////////////////////////////////////////////////////////
	homes := repository.Homestay{}
	repo.db.First(&homes, input.HomestayID)
	////////////////////////////////////////proses update data avgrating homestay ketika comment ditambahkan/////////////////////////////////////////////////////////////////////////////////////////////
	home := repository.Homestay{}
	home.AvgRate = (homes.AvgRate + float64(input.Ratings)) / float64((jumData))

	tx2 := repo.db.Model(&home).Where("id = ?", input.HomestayID).Updates(&home)

	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
}

// DeleteById implements comment.RepositoryInterface
func (repo *commentRepository) DeleteById(id int) (comment.CoreComment, error) {
	komen := Comment{}
	tx1 := repo.db.Delete(&komen, id)
	if tx1.Error != nil {
		return comment.CoreComment{}, tx1.Error
	}

	txres := repo.db.Unscoped().Where("id=?", id).Find(&komen)
	if txres.Error != nil {
		return comment.CoreComment{}, txres.Error
	}
	if tx1.RowsAffected == 0 {
		return comment.CoreComment{}, errors.New("id not found")

	}
	result := komen.ModelsToCore()
	return result, nil
}

// GetAllComment implements comment.RepositoryInterface
func (*commentRepository) GetAllComment() (data []comment.CoreComment, err error) {
	panic("unimplemented")
}

// GetById implements comment.RepositoryInterface
func (*commentRepository) GetById(id int) (data comment.CoreComment, err error) {
	panic("unimplemented")
}

// UpdateComment implements comment.RepositoryInterface
func (repo *commentRepository) UpdateComment(id int, input comment.CoreComment) error {
	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	komen1 := Comment{}
	txres := repo.db.Where("id=?", id).Find(&komen1)
	if txres.Error != nil {
		return txres.Error
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	var commentModel []Comment

	repo.db.Where("homestay_id=?", komen1.HomestayID).Find(&commentModel)

	jumData := len(commentModel)
	log.Println("jumData waktu update", jumData)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	homes := repository.Homestay{}
	repo.db.First(&homes, komen1.HomestayID)
	////////////////////////////////////////proses update data avgrating homestay ketika comment ditambahkan/////////////////////////////////////////////////////////////////////////////////////////////
	home := repository.Homestay{}
	home.AvgRate = ((homes.AvgRate * float64(jumData)) - float64(komen1.Ratings) + float64(input.Ratings)) / float64(jumData)

	tx2 := repo.db.Model(&home).Where("id = ?", komen1.HomestayID).Updates(&home)

	if tx2.Error != nil {
		return tx2.Error
	}
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	komen := FromCore(input)
	input.HomestayID = komen1.HomestayID

	tx := repo.db.Model(&komen).Where("id = ?", id).Updates(&komen)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
