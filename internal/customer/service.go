package customer

func SaveCustomer(customer CustomerRequest) (Customer, error) {
  newCustomer := NewCustomer(customer)

  res, err := Save(newCustomer)
  return res, err
}

func FindAllCustomers() ([]CustomerResponse, error) {
  res, err := FindAll()
  if err != nil {
    return nil, err
  }

  var customers []CustomerResponse

  for _, person := range res {
    customer := NewCustomerResponse(person)

    customers = append(customers, customer)
  }

  return customers, nil
}

func FindById(id string) (CustomerResponse, error) {
  res, err := FindOne(id)
  if err != nil {
    return CustomerResponse{}, err
  }

  customer := NewCustomerResponse(res)
  return customer, nil
}
