package transaction

import (
	"errors"
	"strconv"

	"github.com/Abdirahman04/bytebank-gin/internal/account"
	"github.com/Abdirahman04/bytebank-gin/validations"
)

func ChangeAmount(transaction Transaction) error {
  id := strconv.FormatUint(uint64(transaction.AccountId), 10)
  typ := transaction.TransactionType

  var err error

  if typ == "deposit" {
    err = account.ChandeAccountAmount(id, transaction.Balance)
  } else if typ == "withdraw" {
    err = account.ChandeAccountAmount(id, -transaction.Balance)
  } else {
    target := transaction.Target
    err = account.ChandeAccountAmount(id, -transaction.Balance)
    if err != nil {
      return err
    }

    err = account.ChandeAccountAmount(target, transaction.Balance)
  }

  return err
}

func SaveTransaction(transaction TransactionRequest) (TransactionResponse, error) {
  validType := validations.CheckTransactionType(transaction.TransactionType)
  if !validType {
    return TransactionResponse{}, errors.New("invalid transaction type")
  }

  if transaction.TransactionType == "transfer" {
    _, err := account.FindOne(transaction.Target)
    if err != nil {
      return TransactionResponse{}, errors.New("invalid target id")
    }
  }

  _, err := account.FindOne(strconv.FormatUint(uint64(transaction.AccountId), 10))
  if err != nil {
    return TransactionResponse{}, errors.New("invalid account id")
  }

  newTransaction := NewTransaction(transaction)

  res, err := Save(newTransaction)
  if err != nil {
    return TransactionResponse{}, err
  }

  err = ChangeAmount(res)
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

func FindTransactionById(id string) (TransactionResponse, error) {
  res, err := FindById(id)
  if err != nil {
    return TransactionResponse{}, err
  }

  transaction := NewTransactionResponse(res)

  return transaction, nil
}

func FindTransactionsByAccountId(id string) ([]TransactionResponse, error) {
  res, err := FindAllById(id)
  if err != nil {
    return nil, err
  }

  var transactions []TransactionResponse

  for _, trans := range res {
    transaction := NewTransactionResponse(trans)

    transactions = append(transactions, transaction)
  }

  return transactions, nil
}

func DeleteTransaction(id string) {
  Delete(id)
}
