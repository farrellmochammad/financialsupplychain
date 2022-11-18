package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./signed_db.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./signed_db.db")
	if errDB != nil {
		panic(errDB)
	}

	signed_table := `CREATE TABLE signed (
		"SignId" TEXT NOT NULL PRIMARY KEY,
		"FundId" TEXT NOT NULL,
		"SignUrl" TEXT);`

	query, err := db.Prepare(signed_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	db.Close()
	fmt.Println("Signed Table created successfully!")
}
