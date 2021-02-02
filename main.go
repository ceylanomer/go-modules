package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ceylanomer/go-modules/models"
	"github.com/ceylanomer/go-modules/services"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	InitializeDB()
	StartAPI()
}

//InitializeDB is initializes the db
func InitializeDB() {
	connString := "sqlserver://sa:Sa123456@localhost:1433?database=goDB"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Customer{})
	db.Create(&models.Customer{
		Name:        "TestCustomer",
		Address:     "Istanbul",
		PhoneNumber: "5425425454",
	})

	var customer models.Customer

	db.First(&customer, 1)
	db.First(&customer, "Address = ?", "Istanbul")
	db.Model(&customer).Update("Address", "Atasehir")

}

//StartAPI Starts the API
func StartAPI() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/customers", services.GetAllCustomers).Methods("GET")
	fmt.Println("Now listening on: localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
