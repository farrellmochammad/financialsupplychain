package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./pond_db.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./pond_db.db")
	if errDB != nil {
		panic(errDB)
	}

	pond_table := `CREATE TABLE pond (
	    "PondId" TEXT NOT NULL PRIMARY KEY,
		"Nik" TEXT NOT NULL,
	    "TotalSpawning" INTEGER,
		"FishType" TEXT);`

	query, err := db.Prepare(pond_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	db.Close()
	fmt.Println("Pond Table created successfully!")
}
