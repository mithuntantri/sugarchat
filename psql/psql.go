package psql

import (
  "os"
  "log"
  "database/sql"
  _ "github.com/lib/pq"
  "sugarchat/store"
)

var user_name = os.Getenv("POSTGRESQL_USERNAME")
var db_name = "sugarchat"
var password = os.Getenv("POSTGRESQL_PASSWORD")

func ConnectPSQL()  {
  db_conn, db_err := sql.Open(
    "postgres",
    "user=" + user_name + " dbname=" + db_name + " sslmode=disable password=" + password,
  )
  if db_err != nil{
    store.Error.Println("PSQL Connection to " + db_name + " failed")
    log.Fatal(db_err)
  } else {
    store.Success.Println("PSQL Connection to " + db_name + "  Established")
  }
  store.DBConn = db_conn
  CreateUser()
}

func CreateUser(){
  stmt, err := store.DBConn.Prepare(`
      CREATE TABLE IF NOT EXISTS "users" (
        "mobile_number" varchar(10) NOT NULL,
        "email_id" varchar(127) NOT NULL,
        "firstname" varchar(30) NOT NULL,
        "lastname" varchar(30),
        "dob" bigint NOT NULL,
        "gender_code" int NOT NULL,
        "password" varchar(255) NOT NULL,
        CONSTRAINT user_pk PRIMARY KEY ("mobile_number")
      ) WITH (
        OIDS=TRUE
      )
    `)
  if err != nil {
    store.Error.Println("Failed to create Users Table")
    log.Fatal(err)
  } else {
    store.Success.Println("Creating Users Table successful")
  }
  defer stmt.Close()
  if _, err := stmt.Exec(); err != nil{
    log.Fatal(err)
  }
}
