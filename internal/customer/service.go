package customer

func SaveCustomer() (string, error) {
  customer := Customer{FirstName: "Abdirahman", LastName: "Hassan", Email: "abdix@gmail", Password: "killmonger"}

  res, err := Save(customer)
  return res, err
}
