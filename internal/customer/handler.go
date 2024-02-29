package customer

import "github.com/gin-gonic/gin"

func Post(c *gin.Context) {
  var body CustomerRequest

  c.Bind(&body)

  res, err := SaveCustomer(body)
  if err != nil {
    c.Status(400)
    return
  }
  
  c.JSON(200, gin.H{
    "customer": res,
  })
}
