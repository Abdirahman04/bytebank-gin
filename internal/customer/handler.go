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

func GetAll(c *gin.Context) {
  res, err := FindAllCustomers()
  if err != nil {
    c.Status(400)
    return
  }

  c.JSON(200, gin.H{
    "customers": res,
  })
}

func GetOne(c *gin.Context) {
  id := c.Param("id")

  res, err := FindById(id)
  if err != nil {
    c.JSON(400, gin.H{
      "error": "no user found",
    })
    return
  }

  c.JSON(200, gin.H{
    "customer": res,
  })
}
