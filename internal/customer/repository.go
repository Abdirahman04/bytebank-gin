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
