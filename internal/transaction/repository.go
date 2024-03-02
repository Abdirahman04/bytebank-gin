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

func GetAll() ([]Transaction, error) {
  var transactions []Transaction

  result := db.DB.Find(&transactions)
  if result.Error != nil {
    return nil, errors.New("unable to fetch transactions")
  }

  return transactions, nil
}

func FindById(id string) (Transaction, error) {
  var transaction Transaction

  result := db.DB.First(&transaction, id)
  if result.Error != nil {
    return Transaction{}, errors.New("no transaction found")
  }

  return transaction, nil
}
