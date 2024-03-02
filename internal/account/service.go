package account

import (
	"errors"
	"strconv"

	"github.com/Abdirahman04/bytebank-gin/internal/customer"
)

func SaveAccount(account AccountRequest) (AccountResponse, error) {
  _, err := customer.FindById(strconv.Itoa(int(account.CustomerId)))
  if err != nil {
    return AccountResponse{}, errors.New("customerId not found")
  }

  newAccount := NewAccount(account)

  res, err := Save(newAccount)
  if err != nil {
    return AccountResponse{}, err
  }

  resAccount := NewAccountResponse(res)

  return resAccount, nil
}

func FindAllAccounts() ([]AccountResponse, error) {
  res, err := FindAll()
  if err != nil {
    return nil, err
  }

  var accounts []AccountResponse

  for _, account := range res {
    newAccount := NewAccountResponse(account)

    accounts = append(accounts, newAccount)
  }

  return accounts, nil
}

func FindById(id string) (AccountResponse, error) {
  res, err := FindOne(id)
  if err != nil {
    return AccountResponse{}, err
  }

  account := NewAccountResponse(res)

  return account, nil
}
