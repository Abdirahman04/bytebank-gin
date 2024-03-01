package account

import "gorm.io/gorm"

type Account struct {
  gorm.Model
  CustomerId uint
  AccountType uint
  Amount float32
}
