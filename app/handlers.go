package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/golang-concept-microservice/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Sakura endo", "Tokyo", "12345"},
	// 	{"aruno nakanishi", "Osaka", "54321"},
	// }
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Accept") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
		return
	}
	// w.Header().Set("Content-Type", "application/xml")
	// json.NewEncoder(w).Encode(customers)
	// xml.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Customer not found")
		return
	} else {
		if r.Header.Get("Accept") == "application/xml" {
			w.Header().Set("Content-Type", "application/xml")
			xml.NewEncoder(w).Encode(customer)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

}
