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

func UpdateOne(c *gin.Context) {
  id := c.Param("id")

  var body CustomerRequest
  c.Bind(&body)

  res, err := UpdateCustomer(id, body)
  if err != nil {
    c.JSON(400, gin.H{
      "error": err,
    })
    return
  }

  c.JSON(200, gin.H{
    "customer": res,
  })
}

func DeleteCustomer(c *gin.Context) {
  id := c.Param("id")

  DeleteOne(id)

  c.JSON(200, gin.H{
    "message": "account deleted successfully",
  })
}
