package account

import (
	"errors"

	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Save(account Account) (Account, error) {
  result := db.DB.Create(&account)
  if result.Error != nil {
    return Account{}, errors.New("unable to save account")
  }

  return account, nil
}

func FindAll() ([]Account, error) {
  var accounts []Account

  result := db.DB.Find(&accounts)
  if result.Error != nil {
    return nil, result.Error
  }

  return accounts, nil
}

func FindOne(id string) (Account, error) {
  var account Account

  result := db.DB.First(&account, id)
  if result.Error != nil {
    return Account{}, errors.New("no account found")
  }

  return account, nil
}
