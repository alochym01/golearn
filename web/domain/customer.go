// Define Customer business logic Object
package domain

// Customer ...
type Customer struct {
	ID      string
	Name    string
	City    string
	Zipcode string
	Status  string
}

// Data Transfer Object start
// CustomerRequest ...
type CustomerRequest struct {
	Name    string
	City    string
	Zipcode string
}

// CustomerRespone ...
type CustomerRespone struct {
}

// Data Transfer Object send
// CustomerRepository ...
// Define Repository Interface
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	SelectAll() ([]Customer, error)
	Select(id string) (*Customer, error)
	Insert(c Customer) (*Customer, error)
	Update(id string) (*Customer, error)
	Delete(id string) error
}
