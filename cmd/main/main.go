package main

import (
  "log"

  "github.com/Abdirahman04/bytebank-gin/internal/api"
  "github.com/joho/godotenv"
)

func init() {
  err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func main() {
  api.Start()
}
