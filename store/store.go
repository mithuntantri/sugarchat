package store

import (
  "database/sql"
  "github.com/fatih/color"
  r "gopkg.in/dancannon/gorethink.v2"
)

var (
  DBConn *sql.DB
  Session *r.Session
  Statements = color.New(color.FgBlue)
  Info = color.New(color.FgYellow)
  Error = color.New(color.FgRed)
  Success = color.New(color.FgGreen)
  Request = color.New(color.FgCyan)
)
