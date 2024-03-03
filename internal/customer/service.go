package customer

import (
	"github.com/Abdirahman04/bytebank-gin/logger"
	"github.com/Abdirahman04/bytebank-gin/validations"
)

func SaveCustomer(customer CustomerRequest) (Customer, error) {
  al := logger.NewAggregatedLogger()

  customerArr := [4]string{
    customer.FirstName,
    customer.LastName,
    customer.Email,
    customer.Password,
  }

  err := validations.ValidateCustomer(customerArr)
  if err != nil {
    al.Warn(err.Error())
    return Customer{}, err
  }

  newCustomer := NewCustomer(customer)

  res, err := Save(newCustomer)
  return res, err
}

func FindAllCustomers() ([]CustomerResponse, error) {
  al := logger.NewAggregatedLogger()

  res, err := FindAll()
  if err != nil {
    al.Warn(err.Error())
    return nil, err
  }

  var customers []CustomerResponse

  for _, person := range res {
    customer := NewCustomerResponse(person)

    customers = append(customers, customer)
  }

  return customers, nil
}

func FindById(id string) (CustomerResponse, error) {
  al := logger.NewAggregatedLogger()

  res, err := FindOne(id)
  if err != nil {
    al.Warn(err.Error())
    return CustomerResponse{}, err
  }

  customer := NewCustomerResponse(res)
  return customer, nil
}

func UpdateCustomer(id string, customer CustomerRequest) (CustomerResponse, error) {
  al := logger.NewAggregatedLogger()

  newCustomer := NewCustomer(customer)

  res, err := Update(id, newCustomer)
  if err != nil {
    al.Warn(err.Error())
    return CustomerResponse{}, err
  }

  resCustomer := NewCustomerResponse(res)
  
  return resCustomer, nil
}

func DeleteOne(id string) {
  Delete(id)
}
