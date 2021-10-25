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

	// Define CustomerRepositoryAdapterMock Handler
	aMockHandler := storage.NewCustomerRepositoryAdapterMock()
	// Define Service Handler
	svcHandler := service.NewCustomerService(aMockHandler)
	// Define Customer Handler
	cHandler := handler.NewCustomerHandler(svcHandler)

	router.HandleFunc("/customers", cHandler.GetAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
