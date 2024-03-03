package main

import (
	"github.com/Abdirahman04/bytebank-gin/internal/api"
	"github.com/Abdirahman04/bytebank-gin/intializers"
	"github.com/Abdirahman04/bytebank-gin/logger"
	"github.com/Abdirahman04/bytebank-gin/migrate"
	"github.com/Abdirahman04/bytebank-gin/pkg/db"
)

func init() {
  intializers.LoadEnvVariables()
  db.ConnectToDb()
  migrate.Migrate()
  logger.LG = logger.NewAggregatedLogger()
}

func main() {
  api.Start()
}
