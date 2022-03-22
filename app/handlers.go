package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

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
