package transaction

import "github.com/gin-gonic/gin"

func Post(c *gin.Context) {
  var transaction TransactionRequest

  c.Bind(&transaction)

  res, err := SaveTransaction(transaction)
  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "transaction": res,
  })
}

func Get(c *gin.Context) {
  res, err := GetAllTransactions()
  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "transactions": res,
  })
}

func GetById(c *gin.Context) {
  id := c.Param("id")

  res, err := FindTransactionById(id)
  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "transaction": res,
  })
}

func GetByAccountId(c *gin.Context) {
  id := c.Param("id")

  res, err := FindTransactionsByAccountId(id)
  if err != nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "transactions": res,
  })
}

func DeleteOne(c *gin.Context) {
  id := c.Param("id")

  DeleteTransaction(id)

  c.JSON(200, gin.H{
    "message": "transaction deleted successfully",
  })
}
