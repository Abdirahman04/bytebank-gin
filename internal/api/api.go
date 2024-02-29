package api

import (
	"github.com/Abdirahman04/bytebank-gin/internal/customer"
	"github.com/gin-gonic/gin"
)

func Start() {
  r := gin.Default()

  r.GET("/ping", Ping)
  r.POST("/customer", customer.Post)

  r.Run()
}
