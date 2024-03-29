package account

import "github.com/gin-gonic/gin"

func Post(c *gin.Context) {
  var account AccountRequest

  c.Bind(&account)

  res, err := SaveAccount(account)
  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "account": res,
  })
}

func GetAll(c *gin.Context) {
  res, err := FindAllAccounts()
  if err != nil {
    c.JSON(400, gin.H{
      "error": "unable to fetch accounts",
    })
    return
  }

  c.JSON(200, gin.H{
    "accounts": res,
  })
}

func GetOne(c *gin.Context) {
  id := c.Param("id")

  res, err := FindById(id)
  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "account": res,
  })
}

func GetAllById(c *gin.Context) {
  id := c.Param("id")

  res, err := FindAllByCustomerId(id)
  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "accounts": res,
  })
}

func DeleteOne(c *gin.Context) {
  id := c.Param("id")

  DeleteAccount(id)

  c.JSON(200, gin.H{
    "message": "account deleted successfully",
  })
}
