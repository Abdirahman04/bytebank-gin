package customer

import (
	"errors"
	"fmt"

	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(customer Customer) (string, error) {
  result := db.DB.Create(&customer)
  if result.Error != nil {
    return "", errors.New("Error saving customer")
  }

  return fmt.Sprintln(result.RowsAffected), nil
}
