package check

type ServiceInterface interface {
	Checkroom(id int, check_in, check_out string) (string, error)
}

type RepositoryInterface interface {
	Checkroom(id int, check_in, check_out string) (string, error)
}
