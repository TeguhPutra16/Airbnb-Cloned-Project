package service

import (
	"be13/project/features/check"
	"be13/project/features/homestay"
)

type checkService struct {
	checkRepository check.RepositoryInterface
}

func NewCheck(repo check.RepositoryInterface) check.ServiceInterface {
	return &checkService{
		checkRepository: repo,
	}
}

// GetAllhomestay implements check.ServiceInterface
func (*checkService) GetAllhomestay(check_in string, checkout_out string) (data []homestay.CoreHomestay, err error) {
	panic("unimplemented")
}
