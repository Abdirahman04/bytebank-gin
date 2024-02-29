package customer

import "gorm.io/gorm"

type Customer struct {
  gorm.Model
  FirstName string
  LastName string
  Email string
  Password string
}

type CustomerRequest struct {
  FirstName string
  LastName string
  Email string
  Password string
}

func NewCustomer(customer CustomerRequest) Customer {
  return Customer{
    FirstName: customer.FirstName,
    LastName: customer.LastName,
    Email: customer.Email,
    Password: customer.Password,
  }
}
