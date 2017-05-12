package store

import (
    "golang.org/x/crypto/bcrypt"
)

func BcryptPassword(password string) string{
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
        Error.Println("Failed to Hash password")
    }
  return string(hashedPassword)
}

func VerifyBcrypt(db_password, password string) bool{
  err := bcrypt.CompareHashAndPassword([]byte(db_password), []byte(password))
  if err != nil{
    return false
  }
  return true
}
