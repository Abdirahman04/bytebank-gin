package account

import "github.com/gin-gonic/gin"

func Post(c *gin.Context) {
  var account AccountRequest

  c.Bind(&account)

  res, err := SaveAccount(account)
  if err != nil {
    c.JSON(400, gin.H{
      "error": "unable to save account",
    })
    return
  }

  c.JSON(200, gin.H{
    "account": res,
  })
}
