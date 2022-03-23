package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/golang-concept-microservice/domain"
	"github.com/xvbnm48/golang-concept-microservice/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomersRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
