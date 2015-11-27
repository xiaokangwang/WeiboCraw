package main

import "database/sql"
import _ "github.com/mattn/go-sqlite3"

type ExecTimeDbS struct{
dbc * sql.DB
}

func (e *ExecTimeDbS)Boot()error{
  var err error
  db, err := sql.Open("sqlite3", "file:weibopulldb.sqlite")

  if err != nil {
    return err
  }
  defer db.Close()
}
