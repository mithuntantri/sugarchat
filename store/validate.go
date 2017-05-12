package store

import (
  "regexp"
)

func ValidateEmail(email_id string) bool {
 	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
 	return Re.MatchString(email_id)
}

func ValidateMobile(mobile_number string) bool{
  Re := regexp.MustCompile(`\d{10}`)
  return Re.MatchString(mobile_number)
}
