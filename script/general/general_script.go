package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./experience_db.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./experience_db.db")
	if errDB != nil {
		panic(errDB)
	}

	experience_table := `CREATE TABLE experience (
		"Nik" TEXT NOT NULL PRIMARY KEY,
		"Name" TEXT NOT NULL,
		"Phone" TEXT NOT NULL,
		"SubmitBy" TEXT NOT NULL,
		"Dob" TEXT,
		"Address" TEXT,
		"StartFarming" TEXT,
		"FishType" TEXT,
		"NumberOfPonds" INTEGER,
	    "Notes" TEXT,
		"UrlFile" TEXT,
		"FundId" TEXT,
		"CurrentStatus" TEXT);`

	query, err := db.Prepare(experience_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	db.Close()
	fmt.Println("Experience Table created successfully!")

	filefunder, err := os.Create("./funder_db.db")
	if err != nil {
		log.Fatal(err)
	}
	filefunder.Close()

	dbfunder, errDB := sql.Open("sqlite3", "./funder_db.db")
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

	queryfunder, err := dbfunder.Prepare(funder_table)
	if err != nil {
		log.Fatal(err)
	}
	queryfunder.Exec()

	db.Close()
	fmt.Println("Funder Table created successfully!")

	filesigned, err := os.Create("./signed_db.db")
	if err != nil {
		log.Fatal(err)
	}
	filesigned.Close()

	dbsigned, errDB := sql.Open("sqlite3", "./signed_db.db")
	if errDB != nil {
		panic(errDB)
	}

	signed_table := `CREATE TABLE signed (
		"SignId" TEXT NOT NULL PRIMARY KEY,
		"FundId" TEXT NOT NULL,
		"SignUrl" TEXT);`

	querysigned, err := dbsigned.Prepare(signed_table)
	if err != nil {
		log.Fatal(err)
	}
	querysigned.Exec()

	db.Close()
	fmt.Println("Signed Table created successfully!")
}
