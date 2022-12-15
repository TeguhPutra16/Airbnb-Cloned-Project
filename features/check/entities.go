package check

import "be13/project/features/homestay"

type ServiceInterface interface {
	GetAllhomestay(check_in, checkout_out string) (data []homestay.CoreHomestay, err error)
}

type RepositoryInterface interface {
	GetAllhomestay(check_in, checkout_out string) (data []homestay.CoreHomestay, err error)
}
