package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./credit_db.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./credit_db.db")
	if errDB != nil {
		panic(errDB)
	}

	credit_table := `CREATE TABLE credit (
	    "CreditId" TEXT NOT NULL PRIMARY KEY,
	    "CreditAmount" TEXT,
	    "Status" INTEGER);`

	query, err := db.Prepare(credit_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	db.Close()
	fmt.Println("Credit Table created successfully!")
}
