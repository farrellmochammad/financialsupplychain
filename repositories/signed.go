package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

func InsertSigned(sign *__models.Signed) {
	db, errDB := sql.Open("sqlite3", "./signed_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO signed(SignId, FundId, SignUrl) VALUES (?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(NewSHA1Hash(), sign.FundId, sign.SignUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetSignsByFundId(fundid []string) []__models.Signed {
	db, errDB := sql.Open("sqlite3", "./signed_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM signed where FundId = (?) ", fundid)
	if err != nil {
		panic(err)
	}

	var scanner __models.Signed

	var signs []__models.Signed

	for rows.Next() {
		err = rows.Scan(&scanner.SignId, &scanner.FundId, &scanner.SignUrl)

		if err != nil {
			panic(err)
		}

		signs = append(signs, scanner)
	}

	defer rows.Close()

	defer db.Close()

	return signs
}

func GetSigns() []__models.Signed {
	db, errDB := sql.Open("sqlite3", "./signed_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM signed")
	if err != nil {
		panic(err)
	}

	var scanner __models.Signed

	var signs []__models.Signed

	for rows.Next() {
		err = rows.Scan(&scanner.SignId, &scanner.FundId, &scanner.SignUrl)

		if err != nil {
			panic(err)
		}

		signs = append(signs, scanner)

	}

	defer rows.Close()

	defer db.Close()

	return signs
}

func GetSign(fundid string) __models.Signed {
	db, errDB := sql.Open("sqlite3", "./signed_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM signed")
	if err != nil {
		panic(err)
	}

	var scanner __models.Signed

	for rows.Next() {
		err = rows.Scan(&scanner.SignId, &scanner.FundId, &scanner.SignUrl)

		if err != nil {
			panic(err)
		}

		if fundid == scanner.FundId {
			return scanner
		}

	}

	defer rows.Close()

	defer db.Close()

	return __models.Signed{}
}
