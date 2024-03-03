package validations

import (
	"errors"
	"regexp"
)

func ValidateCustomer(customer [4]string) error {
  if len(customer[0]) < 3 {
    return errors.New("First name should be atleast 3 characters long")
  } else if len(customer[1]) < 3 {
    return errors.New("Last name should be atleast 3 characters long")
  } else if !isValidEmail(customer[2]) {
    return errors.New("Invalid email")
  } else if len(customer[3]) < 5 {
    return errors.New("The password should be atleast 5 characters long")
  } else {
    return nil
  }
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}
