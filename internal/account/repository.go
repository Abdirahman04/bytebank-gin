package account

import "github.com/Abdirahman04/bytebank-gin/pkg/db"

func Save(account Account) (Account, error) {
  result := db.DB.Create(&account)
  if result.Error != nil {
    return Account{}, result.Error
  }

  return account, nil
}
