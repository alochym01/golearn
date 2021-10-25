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

// CustomerRepository ...
// Define Repository Interface
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
