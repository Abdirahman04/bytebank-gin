package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() (gorm.DB, error) {
  var err error
  dsn := os.Getenv("DB_URL")

  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Println("Failed to connect to the database:", err)
    return *DB, err
  }
  return *DB, nil
}
