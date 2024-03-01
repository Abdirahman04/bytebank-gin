package account

import (
	"time"

	"gorm.io/gorm"
)

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

type AccountResponse struct {
  CustomerId uint `json:"customerid"`
  AccountType string `json:"accounttype"`
  Amount float32 `json:"amount"`
  CreatedAt time.Time `json:"createdat"`
  UpdatedAt time.Time `json:"updatedat"`
  DeletedAt time.Time `json:"deletedat"`
}
