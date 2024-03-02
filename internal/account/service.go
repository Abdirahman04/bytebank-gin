package account

func SaveAccount(account AccountRequest) (AccountResponse, error) {
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
