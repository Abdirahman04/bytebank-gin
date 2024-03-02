package transaction

import (
	"errors"

	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(transaction Transaction) (Transaction, error) {
  result := db.DB.Create(&transaction)
  if result.Error != nil {
    return Transaction{}, errors.New("unable to save transaction")
  }

  return transaction, nil
}
