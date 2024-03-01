package account

import "gorm.io/gorm"

type Account struct {
  gorm.Model
  CustomerId uint
  AccountType string
  Amount float32
}

type AccountRequest struct {
  CustomerId uint `json:"customerid"`
  AccountType string `json:"accounttype"`
  Amount float32 `json:"amount"`
}
