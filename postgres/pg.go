package main

/*
requires postgres to be up and running and for an appropriate table and columns
also needs to postgres driver to be installed
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
  "time"
)

func main() {
	db, err := sql.Open("postgres", "user=mvp password=secret dbname=secure_prototype_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, type, updated_at FROM questions")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
    var typ string
    var updatedAt time.Time
		err = rows.Scan(&id, &typ, &updatedAt)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(id, typ, updatedAt)
	}
}
