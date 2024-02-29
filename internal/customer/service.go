package customer

func SaveCustomer(customer CustomerRequest) (Customer, error) {
  newCustomer := NewCustomer(customer)

  res, err := Save(newCustomer)
  return res, err
}
