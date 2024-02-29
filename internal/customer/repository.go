package customer

import (
	"errors"

	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(customer Customer) (Customer, error) {
  result := db.DB.Create(&customer)
  if result.Error != nil {
    return Customer{}, errors.New("Error saving customer")
  }

  return customer, nil
}

func FindAll() ([]Customer, error) {
  var customers []Customer

  result := db.DB.Find(&customers)
  if result.Error != nil {
    return nil, result.Error
  }

  return customers, nil
}

func FindOne(id string) (Customer, error) {
  var customer Customer

  result := db.DB.First(&customer, id)
  if result.Error != nil {
    return customer, result.Error
  }

  return customer, nil
}
