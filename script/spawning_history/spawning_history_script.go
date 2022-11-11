package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./spawning_history.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./spawning_history.db")
	if errDB != nil {
		panic(errDB)
	}

	spawning_table := `CREATE TABLE spawning_history (
		"Nik" TEXT NOT NULL,
	    "Date" TEXT,
	    "Amount" INTEGER,
		"FishType" TEXT);`

	query, err := db.Prepare(spawning_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	db.Close()
	fmt.Println("Spawning history Table created successfully!")
}
