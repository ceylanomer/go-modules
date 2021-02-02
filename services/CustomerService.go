package services

import (
	"encoding/json"
	"net/http"

	"github.com/ceylanomer/go-modules/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//GetAllCustomers Gets all customer list
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	connString := "sqlserver://sa:Sa123456@localhost:1433?database=goDB"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var customer models.Customer
	db.First(&customer, 1)
	json.NewEncoder(w).Encode(customer)

}
