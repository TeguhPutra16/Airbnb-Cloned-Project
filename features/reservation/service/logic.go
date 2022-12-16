package service

import (
	"be13/project/features/reservation"
	"time"

	"github.com/go-playground/validator/v10"
)

// bisnis Logic
type reservasiService struct {
	reservasiRepository reservation.RepositoryEntities //data repository dri entities
	validate            *validator.Validate
}

func NewRes(repo reservation.RepositoryEntities) reservation.ServiceEntities { //dengan kembalian user.service
	return &reservasiService{
		reservasiRepository: repo,
		validate:            validator.New(),
	}
}

// Create implements reservation.ServiceEntities
func (service *reservasiService) Create(input reservation.CoreReservation) error {
	err := service.reservasiRepository.Create(input)
	return err

}

// DeleteById implements reservation.ServiceEntities
func (*reservasiService) DeleteById(id int) (reservation.CoreReservation, error) {
	panic("unimplemented")
}

// GetAllreservasi implements reservation.ServiceEntities
func (service *reservasiService) GetAllreservasi(user int) (data []reservation.CoreReservation, err error) {
	data, err = service.reservasiRepository.GetAllreservasi(user)
	return data, err
}

// GetById implements reservation.ServiceEntities
func (*reservasiService) GetById(id int) (data reservation.CoreReservation, err error) {
	panic("unimplemented")
}

// Update implements reservation.ServiceEntities
func (*reservasiService) Update(id int, input reservation.CoreReservation) error {
	panic("unimplemented")
}

func WaktuNginap(start, end time.Time) float64 {
	difference := start.Sub(end)
	days := float64(difference.Hours() / 24)

	return days

}

func ConverStringtoTime(waktu string) time.Time {
	checkIn, errConvtime1 := time.Parse("02/01/2006", waktu)
	if errConvtime1 != nil {
		return time.Time{}
	}
	return checkIn
}
