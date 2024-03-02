package validations

func CheckTransactionType(typ string) bool {
  if typ == "deposit" || typ == "withdraw" || typ == "transafer" {
    return true
  } else {
    return false
  }
}
