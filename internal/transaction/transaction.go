package transaction

import "gorm.io/gorm"

type Transaction struct {
  gorm.Model
  AccountId uint
  TransactionType string
  Balance float32
}
