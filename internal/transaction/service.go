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
