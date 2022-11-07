package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./monitoring_db.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./monitoring_db.db")
	if errDB != nil {
		panic(errDB)
	}

	spawning_table := `CREATE TABLE monitoring (
	    "FundId" TEXT NOT NULL PRIMARY KEY,
		"Nik" TEXT NOT NULL,
	    "Name" TEXT,
	    "TotalSpawning" INTEGER,
		"FishType" TEXT);`

	query, err := db.Prepare(spawning_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	db.Close()
	fmt.Println("Monitoring Table created successfully!")
}
