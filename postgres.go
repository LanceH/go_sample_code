package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=xxx dbname=pqgotest sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select name, id from blah")
  if err != nil {
    fmt.Println(err)
  }
  defer rows.Close()

  for rows.Next() {
    var name string
    var id int64
    rows.Scan(&name, &id)
    fmt.Println(id, name)
  }
}
