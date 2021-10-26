package service

import "alochym/domain"

// CustomerService ...
// Define Service interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	Create(c domain.Customer) (*domain.Customer, error)
	Read(id string) (*domain.Customer, error)
	Update() (*domain.Customer, error)
	Delete() error
}

// DefaultCustomerService ...
// Implementing All business logic functions
// Implementing all methods in CustomerService Interface
type DefaultCustomerService struct {
	// repo is an CustomerRepository Interface
	repo domain.CustomerRepository
}

// GetAllCumtomer ...
// Implementing all methods in DefaultCustomerService Interface
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) Create(c domain.CustomerRequest) (*domain.Customer, error) {
	cus := domain.Customer{}
	cus.ID = "2000"
	cus.City = c.City
	cus.Name = c.Name
	cus.Zipcode = c.Zipcode
	cus.Status = "1"
	return s.repo.Insert(cus)
	// return nil, nil
}
func (s DefaultCustomerService) Read(id string) (*domain.Customer, error) {
	return s.repo.Select(id)
}
func (s DefaultCustomerService) Update() (*domain.Customer, error) {
	return nil, nil
}
func (s DefaultCustomerService) Delete() error {
	return nil
}

// NewCustomerService ...
// Initial NewCustomerService
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
