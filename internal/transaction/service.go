package transaction

import (
	"errors"
	"strconv"

	"github.com/Abdirahman04/bytebank-gin/internal/account"
)

func SaveTransaction(transaction TransactionRequest) (TransactionResponse, error) {
  _, err := account.FindOne(strconv.FormatUint(uint64(transaction.AccountId), 10))
  if err != nil {
    return TransactionResponse{}, errors.New("invalid account id")
  }

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
