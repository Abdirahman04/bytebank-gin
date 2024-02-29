package customer

import "github.com/gin-gonic/gin"

func Post(c *gin.Context) {
  res, err := SaveCustomer()
  if err != nil {
    c.Status(400)
    return
  }
  
  c.JSON(200, gin.H{
    "result": res,
  })
}
