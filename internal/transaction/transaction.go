package transaction

import "gorm.io/gorm"

type Transaction struct {
  gorm.Model
  AccountId uint
  TransactionType string
  Balance float32
}

type TransactionRequest struct {
  AccountId uint `json:"accountid"`
  TransactionType string `json:"transactiontype"`
  Balance float32 `json:"balance"`
}
