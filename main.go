package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Sakura endo", "Tokyo", "12345"},
		{"aruno nakanishi", "Osaka", "54321"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
