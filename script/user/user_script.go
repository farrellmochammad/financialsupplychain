package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	file, err := os.Create("./username_db.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, errDB := sql.Open("sqlite3", "./username_db.db")
	if errDB != nil {
		panic(errDB)
	}

	users_table := `CREATE TABLE users (
	    "Username" TEXT NOT NULL PRIMARY KEY,
	    "Password" TEXT,
	    "Role" TEXT);`
	query, err := db.Prepare(users_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	records := `INSERT INTO users(Username, Password, Role) VALUES (?, ?, ?)`
	query, err = db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	resC, errC := query.Exec("sales@gmail.com", "123456", "sales")

	if errC != nil {
		log.Fatal(errC)
	}

	fmt.Println(resC.LastInsertId())

	resA, errA := query.Exec("analyst@gmail.com", "123456", "analyst")

	if errA != nil {
		log.Fatal(errC)
	}

	fmt.Println(resA.LastInsertId())

	resB, errB := query.Exec("funder@gmail.com", "123456", "funder")

	if errB != nil {
		log.Fatal(errC)
	}

	fmt.Println(resB.LastInsertId())

	db.Close()
	fmt.Println("Table created successfully!")
}
