package validations

import (
	"errors"
	"regexp"

	"github.com/Abdirahman04/bytebank-gin/internal/customer"
)

func ValidateCustomer(customer customer.CustomerRequest) error {
  if len(customer.FirstName) < 3 {
    return errors.New("First name should be atleast 3 characters long")
  } else if len(customer.LastName) < 3 {
    return errors.New("Last name should be atleast 3 characters long")
  } else if !isValidEmail(customer.Email) {
    return errors.New("Invalid email")
  } else if len(customer.Password) < 5 {
    return errors.New("The password should be atleast 5 characters long")
  } else {
    return nil
  }
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}
