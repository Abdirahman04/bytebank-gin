package customer

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
  gorm.Model
  FirstName string
  LastName string
  Email string
  Password string
}

type CustomerRequest struct {
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email"`
  Password string `json:"password"`
}

type CustomerResponse struct {
  Id uint `json:"id"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email"`
  Password string `json:"password"`
  CreatedAt time.Time `json:"createdat"`
  UpdatedAt time.Time `json:"updatedat"`
  DeletedAt time.Time `json:"deletedat"`
}

func NewCustomer(customer CustomerRequest) Customer {
  return Customer{
    FirstName: customer.FirstName,
    LastName: customer.LastName,
    Email: customer.Email,
    Password: customer.Password,
  }
}

func NewCustomerResponse(customer Customer) CustomerResponse {
  return CustomerResponse{
    Id: customer.ID,
    FirstName: customer.FirstName,
    LastName: customer.LastName,
    Email: customer.Email,
    Password: customer.Password,
    CreatedAt: customer.CreatedAt,
    UpdatedAt: customer.UpdatedAt,
    DeletedAt: customer.DeletedAt.Time,
  }
}
