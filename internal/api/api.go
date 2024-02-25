package api

import "github.com/gin-gonic/gin"

func Start() {
  r := gin.Default()
  r.GET("/ping", Ping)
  r.Run()
}
