package app

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode int
}

func customers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ernest", City: "Harare", Zipcode: 7441},
		{Name: "XXXX", City: "Harare", Zipcode: 7441},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
