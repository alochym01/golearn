package router

import (
	"alochym/handler"
	"alochym/service"
	"alochym/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// Define CustomerRepositoryAdapterMock Adapter
	aMockAdapter := storage.NewCustomerRepositoryAdapterMock()
	// Define Service Handler
	// if any receiver is a point => & symbol
	svcHandler := service.NewCustomerService(&aMockAdapter)
	// Define Customer Handler
	cHandler := handler.NewCustomerHandler(svcHandler)

	router.HandleFunc("/customers", cHandler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", cHandler.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", cHandler.CreateCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
