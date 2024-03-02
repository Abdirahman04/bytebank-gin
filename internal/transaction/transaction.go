package transaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
  gorm.Model
  AccountId uint
  TransactionType string
  Target string
  Balance float32
}

type TransactionRequest struct {
  AccountId uint `json:"accountid"`
  TransactionType string `json:"transactiontype"`
  Target string `json:"target"`
  Balance float32 `json:"balance"`
}

type TransactionResponse struct {
  Id uint `json:"id"`
  AccountId uint `json:"accountid"`
  TransactionType string `json:"transactiontype"`
  Target string `json:"target"`
  Balance float32 `json:"balance"`
  CreatedAt time.Time `json:"createdat"`
  UpdatedAt time.Time `json:"updatedat"`
  DeletedAt time.Time `json:"deletedat"`
}

func NewTransaction(transaction TransactionRequest) Transaction {
  return Transaction{
    AccountId: transaction.AccountId,
    TransactionType: transaction.TransactionType,
    Target: transaction.Target,
    Balance: transaction.Balance,
  }
}

func NewTransactionResponse(transaction Transaction) TransactionResponse {
  return TransactionResponse{
    Id: transaction.ID,
    AccountId: transaction.AccountId,
    TransactionType: transaction.TransactionType,
    Target: transaction.Target,
    Balance: transaction.Balance,
    CreatedAt: transaction.CreatedAt,
    UpdatedAt: transaction.UpdatedAt,
    DeletedAt: transaction.DeletedAt.Time,
  }
}
