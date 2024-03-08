package validations

func CheckTransactionType(typ string) bool {
  if typ == "deposit" || typ == "withdraw" || typ == "transfer" {
    return true
  } else {
    return false
  }
}
