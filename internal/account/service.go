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
