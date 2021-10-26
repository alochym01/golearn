// Can be used as Port/Adaptor/Storage name
package storage

import (
	"alochym/domain"
	"errors"
	"strings"
)

var c = []domain.Customer{
	{ID: "1000", Name: "Do Nguyen Ha", City: "Ho Chi Minh", Zipcode: "70000", Status: "1"},
	{ID: "1001", Name: "Do Thi Kim Hoa", City: "Ho Chi Minh", Zipcode: "70000", Status: "1"},
	{ID: "1002", Name: "Do Thi Hiep", City: "Ho Chi Minh", Zipcode: "70000", Status: "0"},
	{ID: "1003", Name: "Do Nhat Huy", City: "Ho Chi Minh", Zipcode: "70000", Status: "0"},
	{ID: "1004", Name: "Do Nguyen Chuong", City: "Ho Chi Minh", Zipcode: "70000", Status: "0"},
}

// CustomerRepositoryAdapterMock ...
type CustomerRepositoryAdapterMock struct {
	customers []domain.Customer
}

// FindAll ...
// Implementing all methods in CustomerRepository Interface
func (c *CustomerRepositoryAdapterMock) FindAll() ([]domain.Customer, error) {
	return c.customers, nil
	// return nil, errors.New("alochym")
}

// SelectAll ...
func (c *CustomerRepositoryAdapterMock) SelectAll() ([]domain.Customer, error) {
	return c.customers, nil
}

// Insert ...
func (c *CustomerRepositoryAdapterMock) Insert(cus domain.Customer) (*domain.Customer, error) {
	c.customers = append(c.customers, cus)
	return &cus, nil
}

// Select ...
func (c *CustomerRepositoryAdapterMock) Select(id string) (*domain.Customer, error) {
	for _, v := range c.customers {
		if strings.Contains(v.ID, id) {
			return &v, nil
		}
	}
	return nil, errors.New("Not Found")
}

// Update ...
func (c *CustomerRepositoryAdapterMock) Update(id string) (*domain.Customer, error) {
	return nil, nil
}

// Delete ...
func (c *CustomerRepositoryAdapterMock) Delete(id string) error {
	return nil
}

// NewCustomerRepositoryAdapterMock ...
// Initial NewCustomerRepositoryAdapterMock
func NewCustomerRepositoryAdapterMock() CustomerRepositoryAdapterMock {
	return CustomerRepositoryAdapterMock{c}
}
