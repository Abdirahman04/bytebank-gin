package transaction

import (
	"time"

	"gorm.io/gorm"
)

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

type TransactionResponse struct {
  Id uint `json:"id"`
  AccountId uint `json:"accountid"`
  TransactionType string `json:"transactiontype"`
  Balance float32 `json:"balance"`
  CreatedAt time.Time `json:"createdat"`
  UpdatedAt time.Time `json:"updatedat"`
  DeletedAt time.Time `json:"deletedat"`
}
