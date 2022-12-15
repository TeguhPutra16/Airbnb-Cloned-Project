package service

import (
	"be13/project/features/check"
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
func (service *checkService) Checkroom(id int, check_in string, check_out string) (string, error) {
	data, err := service.checkRepository.Checkroom(id, check_in, check_out)
	return data, err
}
