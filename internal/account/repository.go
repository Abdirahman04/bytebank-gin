package account

import (
	"errors"

	"github.com/Abdirahman04/bytebank-gin/logger"
	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(account Account) (Account, error) {
  result := db.DB.Create(&account)
  if result.Error != nil {
    return Account{}, errors.New("unable to save account")
  }

  logger.LG.Info("account created")
  return account, nil
}

func FindAll() ([]Account, error) {
  var accounts []Account

  result := db.DB.Find(&accounts)
  if result.Error != nil {
    return nil, result.Error
  }

  logger.LG.Info("all accounts fetched")
  return accounts, nil
}

func FindOne(id string) (Account, error) {
  var account Account

  result := db.DB.First(&account, id)
  if result.Error != nil {
    return Account{}, errors.New("no account found")
  }

  logger.LG.Info("one account fetched")
  return account, nil
}

func FindAllById(id string) ([]Account, error) {
  var accounts []Account

  result := db.DB.Where("customer_id = ?", id).Find(&accounts)
  if result.Error != nil {
    return nil, errors.New("unable to fetch accounts")
  }
  
  logger.LG.Info("all accounts per customer id fetched")
  return accounts, nil
}

func ChangeAmount(id string, amount float32) error {
  var account Account

  result := db.DB.First(&account, id)
  if result.Error != nil {
    return errors.New("no account found")
  }

  newAmount := account.Amount + amount

  db.DB.Model(&account).Update("amount", newAmount)
  logger.LG.Info("account amount changed")
  return nil
}

func Delete(id string) {
  db.DB.Delete(&Account{}, id)
  logger.LG.Info("account deleted")
}
