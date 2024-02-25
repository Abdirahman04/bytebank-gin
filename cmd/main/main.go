package main

import (
	"github.com/Abdirahman04/bytebank-gin/internal/api"
	"github.com/Abdirahman04/bytebank-gin/intializers"
	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func init() {
  intializers.LoadEnvVariables()
  db.ConnectToDb()
}

func main() {
  api.Start()
}
