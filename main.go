package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ceylanomer/go-modules/customercontext"
	"github.com/ceylanomer/go-modules/models"
	"github.com/ceylanomer/go-modules/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {
	str, err2 := customercontext.CreateDBIfNotExist()
	if err2 == nil {
		fmt.Println(str)
		db, err := customercontext.ConnectDB()
		if err != nil {
			fmt.Println(err)
		}
		MigrateDB(db)
		InitializeDB(db)
		StartAPI()
	} else {
		fmt.Println(str + err2.Error())
	}
}

//MigrateDB migrates db
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Customer{})
}

//InitializeDB is initializes the db
func InitializeDB(db *gorm.DB) {

	db.Create(&models.Customer{
		Name:        "TestCustomer",
		Address:     "Istanbul",
		PhoneNumber: "5425425454",
	})

	// var customer models.Customer

	// db.First(&customer, 1)
	// db.First(&customer, "Address = ?", "Istanbul")
	// db.Model(&customer).Update("Address", "Atasehir")

}

//StartAPI Starts the API
func StartAPI() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/customers", services.GetAllCustomers).Methods("GET")
	myRouter.HandleFunc("/customer", services.AddCustomer).Methods("POST")
	fmt.Println("Now listening on: localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
