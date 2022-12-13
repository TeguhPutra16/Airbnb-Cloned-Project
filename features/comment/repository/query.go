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
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	var commentModel []Comment

	repo.db.Where("homestay_id=?", input.HomestayID).Find(&commentModel)

	jumData := len(commentModel)
	log.Println("jumData", jumData)
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	homes := repository.Homestay{}
	repo.db.First(&homes, input.HomestayID)

	home := repository.Homestay{}

	home.AvgRate = (homes.AvgRate + input.Ratings) / jumData

	tx2 := repo.db.Model(&home).Where("id = ?", input.HomestayID).Updates(&home)

	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
}

// DeleteById implements comment.RepositoryInterface
func (*commentRepository) DeleteById(id int) (comment.CoreComment, error) {
	panic("unimplemented")
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
func (*commentRepository) UpdateComment(id int, input comment.CoreComment) error {
	panic("unimplemented")
}
