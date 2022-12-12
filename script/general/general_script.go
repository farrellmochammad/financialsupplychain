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
		"ProposedBy" TEXT,
		"ApprovedBy" TEXT,
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
		"ProposedBy" TEXT,
		"ProposedTimestamp" TEXT,
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

	filemonitoring, errmonitoring := os.Create("./monitoring_db.db")
	if errmonitoring != nil {
		log.Fatal(err)
	}
	filemonitoring.Close()

	dbmonitoring, errDBmonitoring := sql.Open("sqlite3", "./monitoring_db.db")
	if errDBmonitoring != nil {
		panic(errDB)
	}

	monitoring_table := `CREATE TABLE monitoring (
	    "FundId" TEXT NOT NULL PRIMARY KEY,
		"Nik" TEXT NOT NULL,
	    "Name" TEXT,
	    "TotalSpawning" INTEGER,
		"FishType" TEXT);`

	querymonitoring, err := dbmonitoring.Prepare(monitoring_table)
	if err != nil {
		log.Fatal(err)
	}
	querymonitoring.Exec()

	db.Close()
	fmt.Println("Monitoring Table created successfully!")

	fileSpawningHistory, errSpawningHistory := os.Create("./spawning_history.db")
	if errSpawningHistory != nil {
		log.Fatal(err)
	}
	fileSpawningHistory.Close()

	dbSpawningHistory, errDBSpawningHistory := sql.Open("sqlite3", "./spawning_history.db")
	if errDBSpawningHistory != nil {
		panic(errDBSpawningHistory)
	}

	spawning_table := `CREATE TABLE spawning_history (
		"Nik" TEXT NOT NULL,
	    "Date" TEXT,
	    "Amount" INTEGER,
		"FishType" TEXT);`

	querySpawningHistory, errSpawningHistory := dbSpawningHistory.Prepare(spawning_table)
	if errSpawningHistory != nil {
		log.Fatal(errSpawningHistory)
	}
	querySpawningHistory.Exec()

	db.Close()
	fmt.Println("Spawning history Table created successfully!")

	filecredit, errcredit := os.Create("./credit_db.db")
	if errcredit != nil {
		log.Fatal(errcredit)
	}
	filecredit.Close()

	dbCredit, errDBCredit := sql.Open("sqlite3", "./credit_db.db")
	if errDBCredit != nil {
		panic(errDBCredit)
	}

	credit_table := `CREATE TABLE credit (
	    "CreditId" TEXT NOT NULL PRIMARY KEY,
		"Username" TEXT NOT NULL,
		"Nik" TEXT NOT NUll,
	    "CreditAmount" TEXT,
	    "Status" INTEGER);`

	queryCredit, errCredit := dbCredit.Prepare(credit_table)
	if errCredit != nil {
		log.Fatal(errCredit)
	}
	queryCredit.Exec()

	db.Close()
	fmt.Println("Credit Table created successfully!")

	fileUsername, errUsername := os.Create("./username_db.db")
	if errUsername != nil {
		log.Fatal(errUsername)
	}
	fileUsername.Close()

	dbUsername, errDBUsername := sql.Open("sqlite3", "./username_db.db")
	if errDBUsername != nil {
		panic(errDBUsername)
	}

	users_table := `CREATE TABLE users (
	    "Username" TEXT NOT NULL PRIMARY KEY,
	    "Password" TEXT,
	    "Role" TEXT);`

	queryUsername, errUsername := dbUsername.Prepare(users_table)
	if errUsername != nil {
		log.Fatal(errUsername)
	}
	queryUsername.Exec()

	records := `INSERT INTO users(Username, Password, Role) VALUES (?, ?, ?)`
	queryUsername, errUsername = dbUsername.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	resC, errC := queryUsername.Exec("sales@gmail.com", "123456", "sales")

	if errC != nil {
		log.Fatal(errC)
	}

	fmt.Println(resC.LastInsertId())

	resA, errA := queryUsername.Exec("analyst@gmail.com", "123456", "analyst")

	if errA != nil {
		log.Fatal(errC)
	}

	fmt.Println(resA.LastInsertId())

	resB, errB := queryUsername.Exec("funder@gmail.com", "123456", "funder")

	if errB != nil {
		log.Fatal(errC)
	}

	fmt.Println(resB.LastInsertId())

	dbUsername.Close()
	fmt.Println("Username Table created successfully!")
}
