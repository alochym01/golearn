package service

import "alochym/domain"

// CustomerService ...
// Define Service interface
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
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

// NewCustomerService ...
// Initial NewCustomerService
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
