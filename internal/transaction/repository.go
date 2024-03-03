package transaction

import (
	"errors"

	"github.com/Abdirahman04/bytebank-gin/logger"
	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(transaction Transaction) (Transaction, error) {
  result := db.DB.Create(&transaction)
  if result.Error != nil {
    return Transaction{}, errors.New("unable to save transaction")
  }

  logger.LG.Info("transaction created")
  return transaction, nil
}

func GetAll() ([]Transaction, error) {
  var transactions []Transaction

  result := db.DB.Find(&transactions)
  if result.Error != nil {
    return nil, errors.New("unable to fetch transactions")
  }

  logger.LG.Info("all transactions fetched")
  return transactions, nil
}

func FindById(id string) (Transaction, error) {
  var transaction Transaction

  result := db.DB.First(&transaction, id)
  if result.Error != nil {
    return Transaction{}, errors.New("no transaction found")
  }

  logger.LG.Info("one transaction fetched")
  return transaction, nil
}

func FindAllById(id string) ([]Transaction, error) {
  var transactions []Transaction

  result := db.DB.Where("account_id = ?", id).Find(&transactions)
  if result.Error != nil {
    return nil, result.Error
  }

  logger.LG.Info("all transactions per account id fetched")
  return transactions, nil
}

func Delete(id string) {
  db.DB.Delete(&Transaction{}, id)
  logger.LG.Info("transaction deleted")
}
