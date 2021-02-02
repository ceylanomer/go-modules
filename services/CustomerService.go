package services

import (
	"encoding/json"
	"net/http"

	"github.com/ceylanomer/go-modules/customercontext"
	"github.com/ceylanomer/go-modules/models"
)

//GetAllCustomers gets all customer list
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	db, err := customercontext.ConnectDB()

	if err != nil {
		panic("failed to connect database")
	}

	var customer models.Customer
	db.First(&customer, 1)

	json.NewEncoder(w).Encode(customer)
}

//AddCustomer inserts customer
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	var p models.Customer

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err2 := customercontext.ConnectDB()
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&p)

	json.NewEncoder(w).Encode(p)
}
