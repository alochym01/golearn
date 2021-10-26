package handler

import (
	"alochym/domain"
	"alochym/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandler ...
type CustomerHandler struct {
	svc service.CustomerService
}

// GetAllCustomers ...
func (h CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Println()
	c, err := h.svc.GetAllCustomer()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
	return
}

// GetCustomer ...
func (h CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// fmt.Println()
	c, err := h.svc.Read(params["id"])
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
	return
}

// CreateCustomer ...
func (h CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// fmt.Println()
	cus := domain.CustomerRequest{}
	c, err := h.svc.Create(cus)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
	return
}

// NewCustomerHandler ...
func NewCustomerHandler(svc service.DefaultCustomerService) CustomerHandler {
	return CustomerHandler{svc}
}
