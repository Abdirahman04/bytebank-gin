package customer

import (
	"errors"
	"log"

	"github.com/Abdirahman04/bytebank-gin/logger"
	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(customer Customer) (Customer, error) {
  result := db.DB.Create(&customer)
  if result.Error != nil {
    return Customer{}, errors.New("Error saving customer")
  }

  logger.LG.Info("customer created")
  return customer, nil
}

func FindAll() ([]Customer, error) {
  var customers []Customer

  result := db.DB.Find(&customers)
  if result.Error != nil {
    return nil, result.Error
  }

  logger.LG.Info("all customers fetched")
  return customers, nil
}

func FindOne(id string) (Customer, error) {
  var customer Customer

  result := db.DB.First(&customer, id)
  if result.Error != nil {
    return customer, result.Error
  }

  logger.LG.Info("one customer fetched")
  return customer, nil
}

func Update(id string, customer Customer) (Customer, error) {
  var oldCustomer Customer

  result := db.DB.First(&oldCustomer, id)
  if result.Error != nil {
    return oldCustomer, result.Error
  }
  log.Println(oldCustomer)

  db.DB.Model(&oldCustomer).Updates(Customer{
    FirstName: customer.FirstName,
    LastName: customer.LastName,
    Email: customer.Email,
    Password: customer.Password,
  })
  log.Println(oldCustomer)

  logger.LG.Info("customer updated")
  return customer, nil
}

func Delete(id string) {
  db.DB.Delete(&Customer{}, id)
  logger.LG.Info("customer deleted")
}
