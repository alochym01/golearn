package storage

import "alochym/domain"

// CustomerRepositoryAdapterMock ...
type CustomerRepositoryAdapterMock struct {
	customers []domain.Customer
}

// FindAll ...
// Implementing all methods in CustomerRepository Interface
func (c CustomerRepositoryAdapterMock) FindAll() ([]domain.Customer, error) {
	return c.customers, nil
	// return nil, errors.New("alochym")
}

// NewCustomerRepositoryAdapterMock ...
// Initial NewCustomerRepositoryAdapterMock
func NewCustomerRepositoryAdapterMock() CustomerRepositoryAdapterMock {
	c := []domain.Customer{
		{ID: "1000", Name: "Do Nguyen Ha", City: "Ho Chi Minh", Zipcode: "70000", Status: "1"},
		{ID: "1001", Name: "Do Thi Kim Hoa", City: "Ho Chi Minh", Zipcode: "70000", Status: "1"},
		{ID: "1002", Name: "Do Thi Hiep", City: "Ho Chi Minh", Zipcode: "70000", Status: "0"},
		{ID: "1003", Name: "Do Nhat Huy", City: "Ho Chi Minh", Zipcode: "70000", Status: "0"},
		{ID: "1004", Name: "Do Nguyen Chuong", City: "Ho Chi Minh", Zipcode: "70000", Status: "0"},
	}
	return CustomerRepositoryAdapterMock{c}

}
