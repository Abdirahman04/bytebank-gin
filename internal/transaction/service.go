package transaction

func SaveTransaction(transaction TransactionRequest) (TransactionResponse, error) {
  newTransaction := NewTransaction(transaction)

  res, err := Save(newTransaction)
  if err != nil {
    return TransactionResponse{}, err
  }

  resTransaction := NewTransactionResponse(res)
  return resTransaction, nil
}

func GetAllTransactions() ([]TransactionResponse, error) {
  var transactions []TransactionResponse

  res, err := GetAll()
  if err != nil {
    return nil, err
  }

  for _, transaction := range res {
    newTransaction := NewTransactionResponse(transaction)

    transactions = append(transactions, newTransaction)
  }

  return transactions, nil
}
