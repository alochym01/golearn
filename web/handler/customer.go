package handler

import (
	"alochym/service"
	"encoding/json"
	"net/http"
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

// NewCustomerHandler ...
func NewCustomerHandler(svc service.DefaultCustomerService) CustomerHandler {
	return CustomerHandler{svc}
}
