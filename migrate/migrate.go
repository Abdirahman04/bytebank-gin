package migrate

import (
	"github.com/Abdirahman04/bytebank-gin/internal/account"
	"github.com/Abdirahman04/bytebank-gin/internal/customer"
	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func Migrate() {
  db.DB.AutoMigrate(&customer.Customer{})
  db.DB.AutoMigrate(&account.Account{})
}
