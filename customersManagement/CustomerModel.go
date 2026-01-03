package main

type CustomerModel struct {
	Id     uint
	Name   string //
	Gender string
	Email  string
	Age    uint8
	Phone  string
}

func NewCustomerModel(name, email, phone, gender string, age uint8) CustomerModel {
	return CustomerModel{
		Name:   name,
		Gender: gender,
		Email:  email,
		Age:    age,
		Phone:  phone,
	}
}
