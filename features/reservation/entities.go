package reservation

import (
	"time"

	"gorm.io/gorm"
)

type CoreReservation struct {
	gorm.Model
	ID         uint
	HomestayID uint
	UserID     uint
	CheckIn    string
	CheckOut   string
	Price      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
