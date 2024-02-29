package customer

import (
	"errors"
	"fmt"

	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(customer Customer) (string, error) {
  db, err := db.ConnectToDb()
  if err != nil {
    return "", errors.New("Error connecting to db")
  }

  result := db.Create(&customer)
  
  return fmt.Sprintln(result.RowsAffected), nil
}
