package psql

import (
  "sugarchat/store"
)

func CheckUserConflict(key string) bool{
  var count int8
  if store.ValidateEmail(key){
    store.DBConn.QueryRow("SELECT COUNT(*) FROM users WHERE email_id=$1",key).Scan(&count)
  } else {
    store.DBConn.QueryRow("SELECT COUNT(*) FROM users WHERE mobile_number=$1",key).Scan(&count)
  }
  if count == 0 {
    return false
  } else{
    return true
  }
}

func FetchPassword(key string) string{
  var password string
  if store.ValidateEmail(key){
    store.DBConn.QueryRow("SELECT password FROM users WHERE email_id=$1",key).Scan(&password)
  } else {
    store.DBConn.QueryRow("SELECT password FROM users WHERE mobile_number=$1",key).Scan(&password)
  }
  return password
}

func RegisterUser(mobile_number, email_id, firstname, lastname, password string, gender_code, dob int) bool{
  query := `INSERT INTO users(mobile_number, email_id, firstname, lastname, dob, gender_code, password) VALUES($1, $2, $3, $4, $5, $6, $7);`
  stmt, err := store.DBConn.Prepare(query)
  if err != nil {
    store.Error.Println("Failed to Register New User")
    return false
  }
  defer stmt.Close()
  res, err := stmt.Exec(mobile_number, email_id, firstname, lastname, dob, gender_code, password)
  if err != nil || res == nil {
    store.Error.Println("Failed to Register New User")
    return false
  }
  return true
}
