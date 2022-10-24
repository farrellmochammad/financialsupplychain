package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

func InsertCredit(credit *__models.Credit) {
	db, errDB := sql.Open("sqlite3", "./credit_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO credit(CreditId, Username, Nik, CreditAmount, Status) VALUES (?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(credit.CreditId, credit.Username, credit.Nik, credit.CreditAmount, credit.Status)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetCredits() []__models.Credit {
	db, errDB := sql.Open("sqlite3", "./credit_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM credit")
	if err != nil {
		panic(err)
	}

	var scanner __models.Credit

	var credits []__models.Credit

	for rows.Next() {
		err = rows.Scan(&scanner.CreditId, &scanner.Username, &scanner.Nik, &scanner.CreditAmount, &scanner.Status)

		if err != nil {
			panic(err)
		}

		credits = append(credits, scanner)

	}

	defer rows.Close()

	defer db.Close()

	return credits
}

func GetCredit(creditid string) __models.Credit {
	db, errDB := sql.Open("sqlite3", "./credit_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM credit")
	if err != nil {
		panic(err)
	}

	var scanner __models.Credit
	var credit __models.Credit

	for rows.Next() {
		err = rows.Scan(&scanner.CreditId, &scanner.Username, &scanner.Nik, &scanner.CreditAmount, &scanner.Status)

		if err != nil {
			panic(err)
		}

		if creditid == scanner.CreditId {
			credit = scanner
			return credit
		}

	}

	defer rows.Close()

	defer db.Close()

	return credit
}
