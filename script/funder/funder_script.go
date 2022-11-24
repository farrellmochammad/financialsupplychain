package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./funder_db.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	funder_table := `CREATE TABLE funder (
		"FundId" TEXT NOT NULL PRIMARY KEY,
		"Nik" TEXT NOT NULL,
		"SubmittedBy" TEXT NOT NULL,
		"SubmittedTimestamp" TEXT,
		"FundedBy" TEXT,
		"FundedTimestamp" TEXT,
		"FileUrl" TEXT,
		"FishType" TEXT, 
		"NumberOfPonds" INTEGER,
		"AmountOfFund" INTEGER);`

	query, err := db.Prepare(funder_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	db.Close()
	fmt.Println("Funder Table created successfully!")
}
