package rethink

import (
  "sugarchat/store"
  r "gopkg.in/dancannon/gorethink.v2"
)

var (
  db_name = "sugarchat"
)

func ConnectDB()  {
  var err error
  store.Session, err = r.Connect(r.ConnectOpts{
    Address:  "localhost:28015",
    Database: db_name,
    MaxOpen:  40,
  })
  if err != nil {
    store.Error.Println("Connection to RethinkDb Failed")
  } else {
    store.Success.Println("Connection to RethinkDb Successful")
  }
  CreateDB()
  CreateUserSession()
}

func CreateDB()  {
  store.Info.Println("Creating the RethinkDB Database")
  _, err := r.Branch(
    r.DBList().Contains(db_name),
    nil,
    r.DBCreate(db_name),
  ).Run(store.Session)
  if err != nil{
    store.Error.Println("RethinkDB Database creation failed")
  } else {
    store.Success.Println("RethinkDB Database Created")
  }
}
