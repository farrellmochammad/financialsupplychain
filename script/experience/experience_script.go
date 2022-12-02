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
		"ProposeBy" TEXT NOT NULL,
		"ApprovedBy" TEXT NOT NULL,
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
}
