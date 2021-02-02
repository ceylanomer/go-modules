package models

import "gorm.io/gorm"

//Customer the struct of customer entity
type Customer struct {
	gorm.Model
	Name        string
	Address     string
	PhoneNumber string
}

//Customers List of Customer
type Customers []Customer

//NewCustomer is cosntructor of Customer Entity
func NewCustomer(name, address, phoneNumber string) Customer {
	var c Customer
	c.Name = name
	c.Address = address
	c.PhoneNumber = phoneNumber
	return c
}
