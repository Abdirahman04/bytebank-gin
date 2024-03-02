package api

import (
	"github.com/Abdirahman04/bytebank-gin/internal/account"
	"github.com/Abdirahman04/bytebank-gin/internal/customer"
	"github.com/Abdirahman04/bytebank-gin/internal/transaction"
	"github.com/gin-gonic/gin"
)

func Start() {
  r := gin.Default()

  r.GET("/ping", Ping)

  r.POST("/customer", customer.Post)
  r.GET("/customer", customer.GetAll)
  r.GET("/customer/:id", customer.GetOne)
  r.PUT("/customer/:id", customer.UpdateOne)

  r.POST("/account", account.Post)
  r.GET("/account", account.GetAll)
  r.GET("/account/:id", account.GetOne)
  r.GET("/account/customer/:id", account.GetAllById)
  r.DELETE("/account/:id", account.DeleteOne)

  r.POST("/transaction", transaction.Post)
  r.GET("transaction", transaction.Get)

  r.Run()
}
